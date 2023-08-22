var delay = 250;

var scale = 50;

const width = 50, height = 50;

var cameraX1 = 0,
    cameraX2 = Math.ceil(document.documentElement.clientWidth / scale),
    cameraY1 = 0,
    cameraY2 = Math.ceil(document.documentElement.clientHeight / scale);

cameraX2 = (cameraX2 > width? width : cameraX2);
cameraY2 = (cameraY2 > height? height : cameraY2);

const canvas = document.getElementById('canvas');
const context = canvas.getContext('2d');
context.imageSmoothingEnabled = false;

const fillStyle = 'rgb(203, 0, 0)';

const halfPI = Math.PI / 2;

canvas.width = (document.documentElement.clientWidth + scale > width * scale? width * scale : document.documentElement.clientWidth + scale);
canvas.height = (document.documentElement.clientHeight + scale > height * scale? height * scale : document.documentElement.clientHeight + scale);

context.strokeWidth = Math.round(scale / 10);
context.strokeStyle = '#888';

var images = {};
images[NONE]      = null;
images[AND]       = new Image(); images[AND].src       = 'imgs/and.png';
images[ANGLED]    = new Image(); images[ANGLED].src    = 'imgs/angled.png';
images[BLOCK]     = new Image(); images[BLOCK].src     = 'imgs/block.png';
images[CROSS]     = new Image(); images[CROSS].src     = 'imgs/cross.png';
images[FLASH]     = new Image(); images[FLASH].src     = 'imgs/flash.png';
images[FRWD_SIDE] = new Image(); images[FRWD_SIDE].src = 'imgs/frwdside.png';
images[GET]       = new Image(); images[GET].src       = 'imgs/get.png';
images[MEM_CELL]  = new Image(); images[MEM_CELL].src  = 'imgs/memcell.png';
images[NOT]       = new Image(); images[NOT].src       = 'imgs/not.png';
images[SOURCE]    = new Image(); images[SOURCE].src    = 'imgs/source.png';
images[WIRE]      = new Image(); images[WIRE].src      = 'imgs/wire.png';
images[XOR]       = new Image(); images[XOR].src       = 'imgs/xor.png';
images[DOUBLE]    = new Image(); images[DOUBLE].src    = 'imgs/double.png';
images[UNKNOWN]   = new Image(); images[UNKNOWN].src   = 'imgs/unknown.png';

var colors = {};

Object.keys(images).forEach(function(img) {
    const x = 0,
          y = 31;
    if(img != NONE) {
        images[img].onload = function() {
            const index = y * (images[img].width * 4) + x * 4;
            var canvas = document.createElement('canvas');
            canvas.width = images[img].width;
            canvas.height = images[img].height;
            var ctx = canvas.getContext('2d');
            ctx.drawImage(images[img], 0, 0);

            var data = ctx.getImageData(0, 0, images[img].width, images[img].height);
            var red = data.data[index];
            var green = data.data[index+1];
            var blue = data.data[index+2];

            colors[img] = 'rgb('+red+','+green+','+blue+')';
            console.log(colors[img]);

            data.data[index+3] = 0;
            ctx.putImageData(data, 0, 0);
            images[img].src = canvas.toDataURL();
            images[img].onload = function(){};
        }
    }
});


var packs = {};
packs[0] = [
    WIRE,
    FRWD_SIDE,
    ANGLED,
    CROSS,
    DOUBLE,
];

packs[1] = [
    SOURCE,
    FLASH,
    GET,
    BLOCK,
    NONE,
];

packs[2] = [
    MEM_CELL,
    NOT,
    AND,
    XOR,
    NONE,
]

const lastpack = 2;
var pack = 0;
var type = 0;

var dir = 0;

const minDelay = 10, maxDelay = 650, delayChange = 10;

const minScale = 10, maxScale = 100, scaleChange = 5;

const biasX = document.querySelector('canvas').offsetLeft,
      biasY = document.querySelector('canvas').offsetTop;

var lastPreX = 0, lastPreY = 0;

initCore();
setPack();
setInterval(() => mouse_update(), 10)
