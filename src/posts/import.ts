import { Command } from "commander";

export class ImportCommandFactory {
  public static create(postCommand: Command) {
    postCommand.command('import')
      .option('-f, --file-path <filepath>', 'Path of the file to be imported')
      .action(importPost)
  }
}

interface ImportArguments {
  filePath: string
}

function importPost(args: ImportArguments) {
  console.log(args.filePath)
}
