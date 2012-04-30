;(function($){

	$.extend({
		qs : function(str) {
			var kv = "",
				parts = str.slice(str.indexOf('?') + 1).split('&'),
				map = {};

			// this split function is very simple and can't handle multiple values for the same key
			// ie, key1=a&key1=b should return key1 = [a,b], but instead
			// returns key1=b (the last one is the winner)
			for (var idx = 0, max = parts.length; idx < max; idx++) {
				var kv = parts[idx].split('=');
				map[kv[0]] = kv[1];
			}
			return map;
		}
	});

}(window.jQuery));
