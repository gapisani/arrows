function drawCell(x, y, type, dir, lit) {
    context.fillStyle = lit? colors[type] : 'white';
    context.fillRect(x * scale, y * scale, scale, scale);

    if(type != NONE && images[type] !== undefined && images[type] != null) {
        if(dir == 0)
            context.drawImage(images[type], x * scale, y * scale, scale, scale);
        else {
            context.save();

            context.translate(x * scale, y * scale);
            context.rotate(dir * halfPI);

            switch(dir) {
            case EAST:
                context.drawImage(images[type], 0, -scale, scale, scale);
                break;
            case SOUTH:
                context.drawImage(images[type], -scale, -scale, scale, scale);
                break;
            case WEST:
                context.drawImage(images[type], -scale, 0, scale, scale);
                break;
            }

            context.restore();
        }
    }

    context.strokeRect(x * scale, y * scale, scale, scale);
}

function drawPreview(x, y, contains) {
    if(contains) context.globalAlpha = 0.35;
    else {
        context.clearRect(x * scale, y * scale, scale, scale);
        context.globalAlpha = 0.5;
    }

    drawCell(Math.floor(x), Math.floor(y), packs[pack][type], dir, false);

    context.strokeRect(x * scale, y * scale, scale, scale);
    context.globalAlpha = 1;
}
