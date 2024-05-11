const result = await Bun.build({
    entrypoints: ['./src/htmx.js'],
    outdir: './dist',
    target: 'browser',
    format: 'esm',
})

if(result.success) {
    console.log('⚡bun build complete ⚡')
}

if(!result.success) {
    console.error('bun build failed')
    for(const message of result.logs) {
        console.error(message)
    }
}
