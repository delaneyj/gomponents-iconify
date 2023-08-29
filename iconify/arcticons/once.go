package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Once(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 39.06h11.952c.62-1.026 1.416-2.516 2.516-2.516h10.064c1.126 0 2.04 1.627 2.516 2.574l11.952-.116a14.881 14.881 0 0 0-8.806-2.458c5.976-4.596 6.157-8.021 3.773-10.694l-6.29 4.404l-1.17-3.718a4.78 4.78 0 0 0 4.296-6.609c-1.233-2.605-4.428-3.482-7.275-1.893a7.374 7.374 0 0 0-6.21-1.99c-1.82-3.633-5.72-4.143-8.157-2.227s-2.52 5.534.564 7.844c-1.098 1.525-.906 2.592 2.654 6.22l-1.056 2.373l-6.29-4.403c-3.063 3.555-2.096 6.103 3.774 10.693A15.022 15.022 0 0 0 4.5 39.06Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.968 29.625a15.412 15.412 0 0 0 10.064 0c-.13 1.466-1.299 2.977-2.516 4.403h-5.032c-1.396-1.517-2.449-3.004-2.516-4.403Zm3.794-16.349c3.283.053 3.292.057 6.014 1.845l1.892-4.025l-2.959 1.066l-.465-2.774l-1.94 2.036l-1.852-2.542Z"/><circle cx="16.329" cy="17.286" r="1.887" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="31.498" cy="21.869" r="1.635" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}