declare type Language = 'typescript' | 'go';
export interface CommandOptions {
    input: string;
    language: Language;
    output?: string;
    name?: string;
}
declare const json2type: (options: CommandOptions) => Promise<unknown>;
export default json2type;
