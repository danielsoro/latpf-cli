import { Command } from "commander";
import { ImportCommandFactory } from "./import";
import { ListCommandFactory } from "./list";

export class PostsCommandFactory {
  public static create(rootCommand: Command) {
    const posts = rootCommand.command('posts')
    new ImportCommandFactory(posts)
    new ListCommandFactory(posts)
  }
}
