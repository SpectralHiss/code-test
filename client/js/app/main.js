define(['app/resize','app/copypaste','app/delay'], function(resize,copypaste,delay) {
	resize.registerHooks();
	copypaste.registerHooks();
	delay.registerHooks();
});