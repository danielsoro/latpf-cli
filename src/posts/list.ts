import { Command } from "commander";

export class ListCommandFactory {
  public static create(postsCommand: Command) {
    postsCommand.command('list')
  }
}
