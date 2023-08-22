function canvasMouseDown(e) {
    if(((e.button == 0 || e.button == 2) && e.shiftKey) || e.button == 1) {
        var startX = e.pageX, subX;
        var startY = e.pageY, subY;
    }
    else if(e.button == 0 && Math.floor(e.pageX / scale) < width && Math.floor(e.pageY / scale) < height) {
        var x = e.pageX - biasX, y = e.pageY - biasY;
        SetCell(Math.floor(x / scale) + cameraX1, Math.floor(y / scale) + cameraY1, packs[pack][type], dir);
        drawCell(Math.floor(x / scale), Math.floor(y / scale), packs[pack][type], dir, false);
        canvas.onmousemove = function(e) {
            x = e.pageX - biasX, y = e.pageY - biasY;
            SetCell(Math.floor(x / scale) + cameraX1, Math.floor(y / scale) + cameraY1, packs[pack][type], dir);
            drawCell(Math.floor(x / scale), Math.floor(y / scale), packs[pack][type], dir, false);
        }
    }
    else if(e.button == 1 && Math.floor(e.pageX / scale) < width && Math.floor(e.pageY / scale) < height) {
        var celltype = GetCell(Math.floor((e.pageX - biasX) / scale), Math.floor((e.pageY - biasY) / scale)).type;
        type = celltype
    }
    else if(e.button == 2 && Math.floor(e.pageX / scale) < width && Math.floor(e.pageY / scale) < height) {
        var x = e.pageX - biasX, y = e.pageY - biasY;
        SetCell(Math.floor(x / scale) + cameraX1, Math.floor(y / scale) + cameraY1, NONE, NORTH);
        drawCell(Math.floor(x / scale), Math.floor(y / scale), NONE, 0, false);
        canvas.onmousemove = function(e) {
            x = e.pageX - biasX, y = e.pageY - biasY;
            SetCell(Math.floor(x / scale) + cameraX1, Math.floor(y / scale) + cameraY1, 0, 0);
            drawCell(Math.floor(x / scale), Math.floor(y / scale), NONE, 0, false);
        }
    }
}

function mouseWheel(e) {
    if(e.shiftKey) {
        scale = e.deltaY > 0? (scale <= minScale? minScale : scale - scaleChange) : (scale >= maxScale? maxScale : scale + scaleChange);

        canvas.width = (document.documentElement.clientWidth + scale > width * scale? width * scale : document.documentElement.clientWidth + scale);
        canvas.height = (document.documentElement.clientHeight + scale > height * scale? height * scale : document.documentElement.clientHeight + scale);

        context.strokeWidth = Math.round(scale / 10);
        context.strokeStyle = '#888';

        cameraX2 = Math.ceil(document.documentElement.clientWidth / scale)
        cameraY2 = Math.ceil(document.documentElement.clientHeight / scale);
    }
    else if(e.ctrlKey) {
        delay = e.deltaY > 0? (delay <= minDelay? minDelay : delay - delayChange) : (delay >= maxDelay? maxDelay : delay + delayChange);
        document.getElementById('displaySpeed').innerText = delay + 'ms';
    }
    else {
        pack = e.deltaY > 0? (pack <= 0? lastpack : pack - 1) : (pack >= lastpack? 0 : pack + 1);
        setPack();
    }
}

function keyDown(e) {
    switch(e.code) {
        case 'Numpad1':
        case 'Digit1':
            type = 0;
            break;
        case 'Numpad2':
        case 'Digit2':
            type = 1;
            break;
        case 'Numpad3':
        case 'Digit3':
            type = 2;
            break;
        case 'Numpad4':
        case 'Digit4':
            type = 3;
            break;
        case 'Numpad5':
        case 'Digit5':
            type = 4;
            break;
        case 'ArrowUp': dir = NORTH; break;
        case 'ArrowLeft': dir = WEST; break;
        case 'ArrowDown': dir = SOUTH; break;
        case 'ArrowRight': dir = EAST; break;
        case 'KeyQ': dir = dir <= 0? 3 : dir - 1; break;
        case 'KeyE': dir = dir >= 3? 0 : dir + 1; break;
        case 'KeyW':
            cameraY1 = cameraY2 <= Math.ceil(document.documentElement.clientHeight / scale)? 0 : cameraY1 - 1;
            cameraY2 = cameraY2 <= Math.ceil(document.documentElement.clientHeight / scale)? Math.ceil(document.documentElement.clientHeight / scale) : cameraY2 - 1;
            document.getElementById('displayY').innerText = 'Y: ' + cameraY1;
            update_canvas();
            break;
        case 'KeyA':
            cameraX1 = cameraX2 <= Math.ceil(document.documentElement.clientWidth / scale)? 0 : cameraX1 - 1;
            cameraX2 = cameraX2 <= Math.ceil(document.documentElement.clientWidth / scale)? Math.ceil(document.documentElement.clientWidth / scale) : cameraX2 - 1;
            document.getElementById('displayX').innerText = 'X: ' + cameraX1;
            update_canvas();
            break;
        case 'KeyS':
            cameraY1 = cameraY2 >= height? cameraY1 : cameraY1 + 1;
            cameraY2 = cameraY2 >= height? height : cameraY2 + 1;
            document.getElementById('displayY').innerText = 'Y: ' + cameraY1;
            update_canvas();
                break;
        case 'KeyD':
            cameraX1 = cameraX2 >= width? cameraX1 : cameraX1 + 1;
            cameraX2 = cameraX2 >= width? width : cameraX2 + 1;
            document.getElementById('displayX').innerText = 'X: ' + cameraX1;
            update_canvas();
            break;
        case 'Minus':
            delay = delay <= minDelay? minDelay : delay - delayChange;
            document.getElementById('displaySpeed').innerHTML = `<p>${delay}ms</p>`;
            break;
        case 'Equal':
            delay = delay >= maxDelay? maxDelay : delay + delayChange;
            document.getElementById('displaySpeed').innerHTML = `<p>${delay}ms</p>`;
            break;
        case 'KeyX':
            pack = pack >= lastpack? 0 : pack + 1;
            setPack();
            break;
        case 'KeyZ':
            pack = pack <= 0? lastpack : pack - 1;
            setPack();
            break;
    }
}

function setPack() {
    for(var i = 1; i <= 5; i++) {
        document.querySelector('#display'+i.toString()).src =
            packs[pack][i-1] == NONE? images[UNKNOWN].src : images[packs[pack][i-1]].src;
    }
}

function setPreview(e, x, y) {
    if(e != null) x = Math.floor((e.pageX - biasX) / scale), y = Math.floor((e.pageY - biasY) / scale);
    var w = width? Math.ceil(document.documentElement.clientWidth / scale) : width
    var h = height? Math.ceil(document.documentElement.clientHeight / scale) : height
    if(x >= 0 && y >= 0 &&
        x < Math.ceil(document.documentElement.clientWidth / scale) < w &&
        y < Math.ceil(document.documentElement.clientHeight / scale) < h) {
        const cell = GetCell(lastPreX + cameraX1, lastPreY + cameraY1);
        if(cell != null) {
            drawCell(lastPreX, lastPreY, cell.type, cell.dir, cell.powered);
        }

        drawPreview(x, y, (GetCell(lastPreX + cameraX1, lastPreY + cameraY1).type != 0));

        lastPreX = x;
        lastPreY = y;
    }
}

function showAbout() {
    alert("Ты думал тут что-то будет?");
}

window.addEventListener('keydown', keyDown);
window.addEventListener('wheel', mouseWheel);
canvas.addEventListener('mouseup', function() { canvas.onmousemove = function(){}; })
canvas.addEventListener('mousemove', function(e) { setPreview(e, null, null); })

//remove document resize on Ctrl+MW
window.addEventListener('wheel', function(event) {
    if (event.ctrlKey) {
      event.preventDefault();
    }
}, { passive: false });
//
window.addEventListener('mouseclick', function(event) {
    if (event.button == 1) {
      event.preventDefault();
    }
}, { passive: false });
//remove context menu
canvas.addEventListener('contextmenu', function(event) {
    event.preventDefault();
});

canvas.onmousemove = function(e) {
    setPreview(images[packs[pack][type]], lastPreX, lastPreY);
}

function mouse_update() {
    drawPreview(lastPreX, lastPreY, (GetCell(lastPreX + cameraX1, lastPreY + cameraY1).type != 0));
}

function update_canvas() {
    var canv_x = canvas.width / scale;
    var canv_y = canvas.height / scale;

    for (let y = 0; y <= canv_y ; y++) {
        for (let x = 0; x <= canv_x ; x++) {
            const cell = GetCell(x + cameraX1, y + cameraY1);
            if(cell != null) {
                drawCell(x, y, cell.type, cell.dir, cell.powered);
            }
        }
    }
}
