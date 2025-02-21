import { Command } from "commander";
import { ImportCommandFactory } from "./import";
import { ListCommandFactory } from "./list";

export class PostsCommandFactory {
  public static create(rootCommand: Command) {
    const posts = rootCommand.command('posts')
    ImportCommandFactory.create(posts)
    ListCommandFactory.create(posts)
  }
}
