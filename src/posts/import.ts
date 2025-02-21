import { Command } from "commander";

export class ImportCommandFactory {
  imports: Command
  constructor(posts: Command) {
    this.imports = posts.command('import')
  }
}

