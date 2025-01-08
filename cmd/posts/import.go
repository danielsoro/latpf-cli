package posts

import (
	"encoding/csv"
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/danielsoro/wordpress-cli/lib/wordpress/client"
	"github.com/dromara/carbon/v2"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
	"io"
	"log/slog"
	"os"
	"strconv"
)

type ImportPost struct {
	Command *cobra.Command
}

func printRecords(records []string) {
	for _, record := range records {
		fmt.Println(record)
	}
}

func NewImportCommand(clientType client.WordPressClientType) CreatePost {
	importPostCommand := CreatePost{
		Command: &cobra.Command{
			Use:   "import",
			Short: "Import ports from CSV",
			RunE: func(cmd *cobra.Command, args []string) error {
				wordPressClient, err := client.NewWordPressClient(clientType)
				if err != nil {
					return err
				}

				path, err := cmd.Flags().GetString("file")
				if err != nil {
					return err
				}

				var public [][]string
				switch cmd.Flag("type").Value.String() {
				case "public":
					public = readPublic(path)
					bar := progressbar.Default(int64(len(public)))
					t := table.New().
						Border(lipgloss.NormalBorder()).
						BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
						Headers("N", "Title", "Content", "Link", "Data")

					for index, record := range public {
						if index == 0 || len(record[0]) <= 0 {
							continue
						}

						date := carbon.Parse(record[4]).ToRfc3339MicroString()
						post, err := wordPressClient.CreatePost(record[0], record[1], "publish", date)
						if err != nil {
							return err
						}

						var rows [][]string

						err = bar.Add(1)
						if err != nil {
							return err
						}
						rows = append(rows, []string{
							strconv.Itoa(index), post.Title.Raw, post.Content.Raw[:20], post.Link, post.Date,
						})

						t.Rows(rows...)
					}

					fmt.Println(t)

				case "article":
				case "premium":
				default:
					return fmt.Errorf("invalid type. the post's type accept: [ \"public\", \"article\", \"premium\" ]")
				}

				return nil
			},
		},
	}

	importPostCommand.Command.Flags().StringP("file", "f", "", "The CSV file with posts.")
	importPostCommand.Command.Flags().StringP("type", "t", "", "The post's type. Accept: public, article, premium.")
	importPostCommand.Command.Flags().StringP("status", "s", "publish", "The post's status.")

	_ = importPostCommand.Command.MarkFlagRequired("file")
	_ = importPostCommand.Command.MarkFlagRequired("type")

	return importPostCommand
}

func readPublic(path string) [][]string {
	var records [][]string
	file, err := os.Open(path)
	if err != nil {
		slog.Error("Error opening file", slog.String("error", err.Error()))
	}

	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
		records = append(records, record)
	}

	return records
}
