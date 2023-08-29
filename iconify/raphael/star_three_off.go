package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarThreeOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28.63 12.36a1.996 1.996 0 0 0-1.902-1.383h-.015l-7.15.056l-2.155-6.816a2 2 0 0 0-3.812 0l-2.157 6.816l-7.15-.056h-.017a1.999 1.999 0 0 0-1.164 3.627l5.814 4.158l-2.26 6.783c-.276.828.017 1.74.723 2.25c.35.256.763.384 1.175.384c.418 0 .834-.132 1.19-.392l5.75-4.247l5.75 4.248a1.997 1.997 0 0 0 2.367.008a2.003 2.003 0 0 0 .72-2.25l-2.263-6.783l5.815-4.158a2 2 0 0 0 .74-2.246zm-8.918 5.636l2.73 8.184l-6.94-5.125L8.56 26.18l2.73-8.184l-7.02-5.018l8.627.066L15.5 4.82l2.603 8.225l8.627-.066l-7.018 5.016z"/>`),
		g.Group(children),
	)
}