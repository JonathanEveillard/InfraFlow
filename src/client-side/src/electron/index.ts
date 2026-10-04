import electron from "electron";
import { fileURLToPath } from 'url';
import { dirname } from 'path';

const {app, BrowserWindow} = electron;

// Not included in ES Module, defined __dirname
const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);



console.log(__dirname);
console.log(__filename);

app.on('ready',()=>{

    const window = new BrowserWindow({width:600, height:800});
    window.setMenuBarVisibility(false);

    window.loadFile(path.join(__dirname,'index.html'));
});