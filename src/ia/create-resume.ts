import { Command } from "commander";

export class CreateResumeCommandFactory {
  public static create(iaCommand: Command) {
    iaCommand.command('create-resume')
  }
}
