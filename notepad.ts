import * as fs from 'fs/promises';
import * as readline from 'readline';

// Prompt the user to enter their text
const text = await new Promise<string>((resolve) => {
  const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
  });

  rl.question('Enter your text: ', (answer) => {
    rl.close();
    resolve(answer);
  });
});

// Save the text to a file
await fs.writeFile('output.txt', text);

// Exit the process
process.exit(0);
