import { Command } from "commander";
import { CreateResumeCommandFactory } from "./create-resume";

export class IaCommandFactory {
  public static create(rootCommand: Command) {
    const iaCommand = rootCommand.command('ia')
    CreateResumeCommandFactory.create(iaCommand)
  }
}
