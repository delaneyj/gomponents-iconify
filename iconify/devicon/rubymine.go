package devicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rubymine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<linearGradient id="deviconRubymine0" x1="65.05" x2="52.91" y1="60.03" y2="28.18" gradientUnits="userSpaceOnUse"><stop offset="0" stop-color="#fe2857"/><stop offset=".06" stop-color="#fe3052"/><stop offset=".33" stop-color="#fd533b"/><stop offset=".58" stop-color="#fc6c2a"/><stop offset=".81" stop-color="#fc7b20"/><stop offset="1" stop-color="#fc801d"/></linearGradient><linearGradient id="deviconRubymine1" x1="41.93" x2="60.67" y1="14.45" y2="31.63" gradientUnits="userSpaceOnUse"><stop offset="0" stop-color="#6b57ff"/><stop offset="1" stop-color="#fe2857"/></linearGradient><linearGradient id="deviconRubymine2" x1="3.92" x2="65.63" y1="19.88" y2="98.32" gradientUnits="userSpaceOnUse"><stop offset="0" stop-color="#6b57ff"/><stop offset=".3" stop-color="#fe2857"/><stop offset=".63" stop-color="#fe2857"/><stop offset=".64" stop-color="#fe3052"/><stop offset=".7" stop-color="#fd533b"/><stop offset=".76" stop-color="#fc6c2a"/><stop offset=".81" stop-color="#fc7b20"/><stop offset=".85" stop-color="#fc801d"/></linearGradient><path fill="url(#deviconRubymine0)" d="M101.595 5.486L68.108 17.481L41.716 5.486L33.024 27.38h-5.62v64.634l81.274.707l12.617-64.366z"/><path fill="url(#deviconRubymine1)" d="m100.596 47.482l-53.48-31.695l53.48 62.683z"/><path fill="url(#deviconRubymine2)" d="m52.98 119.467l43.739-5.827l-6.79-13.056h10.667V78.47L47.104 15.689L4.267 26.21l.049 61.44l24.625 34.865l23.906-3.035l.11-.012z"/><path d="M27.429 27.429h73.143v73.143H27.429z"/><path fill="#fff" d="M36.547 86.747h27.429v4.571H36.547zm27.099-50.213h6.522l7.229 11.63l7.229-11.63h6.522v27.49h-6.01V46.08l-7.741 11.739h-.158l-7.656-11.618v17.822h-5.937zm-27.075.037h12.556q5.218 0 7.997 2.779a8.533 8.533 0 0 1 2.353 6.278v.085a8.533 8.533 0 0 1-1.634 5.425a9.435 9.435 0 0 1-4.254 3.084l6.717 9.752h-7.07l-5.693-8.422h-4.986v8.436h-5.986zm12.19 13.3a5.022 5.022 0 0 0 3.377-1.109a3.657 3.657 0 0 0 1.219-2.828v-.073a3.474 3.474 0 0 0-1.219-2.938a5.522 5.522 0 0 0-3.486-.987h-6.071v7.887z"/>`),
		g.Group(children),
	)
}