import { Command } from "commander";
import { PostsCommandFactory } from "./posts";
import { IaCommandFactory } from "./ia";

export function init() {
  const rootCommand = new Command()
    .description('LATPF-CLI to help with some commons tasks for the day')
    .version('0.0.1')

  PostsCommandFactory.create(rootCommand)
  IaCommandFactory.create(rootCommand)
  return rootCommand
}

