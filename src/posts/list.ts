import { Command } from "commander";

export class ListCommandFactory {
  list: Command
  constructor(posts: Command) {
    this.list = posts.command('list')
  }
}
