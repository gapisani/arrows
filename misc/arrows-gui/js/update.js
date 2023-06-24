function initCore() {
    GridInit(width, height);
    Update();
    for(y = 0; y < height; y++) {
        for(x = 0; x < width; x++) {
            cell = GetCell(x, y);
            drawCell(x, y, cell.type, cell.dir, cell.powered);
        }
    }

    updateLoop();

    document.getElementById('load_alert').style.display = 'none';

    canvas.onmousedown = function(e) { canvasMouseDown(e); }
}

async function updateLoop() {
    while(true) {
        var updatePoints = Update();
        var cell;
        for(var i = 0; i < updatePoints.length; i++) {
            var x = updatePoints[i][0],
                y = updatePoints[i][1];
            if(x < cameraX1 || x > cameraX2 || y < cameraY1 || y > cameraY2) { continue }
            cell = GetCell(x, y);
            drawCell(x - cameraX1, y - cameraY1, cell.type, cell.dir, cell.powered);
        }
        await new Promise(resolve => setTimeout(resolve, delay));
    }
}

