import { Command } from "commander";

export class ImportCommandFactory {
  public static create(postCommand: Command) {
    postCommand.command('import')
  }
}

