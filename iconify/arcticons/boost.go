package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Boost(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.42 14.7a4.1 4.1 0 0 0-4.08-4.11h0a4.11 4.11 0 1 0 0 8.21a4.09 4.09 0 0 0 2.88-1.18a4 4 0 0 0 1.19-2.92ZM16.1 29.44l-2-2l-6.2 1.98h-.2a.69.69 0 0 1-.53-.2l-1.45-1.43a.7.7 0 0 1-.11-.88L7.3 24a13 13 0 0 1 10.07-6.39l2.49-.22q2.17-2.58 4-4.42A26.4 26.4 0 0 1 32 7.12a25.63 25.63 0 0 1 9.77-1.61a.79.79 0 0 1 .54.22a.68.68 0 0 1 .22.5A25.27 25.27 0 0 1 40.78 16A25.79 25.79 0 0 1 35 24.16c-1.23 1.24-2.71 2.57-4.42 4l-.22 2.48A13 13 0 0 1 24 40.72l-2.89 1.69a.78.78 0 0 1-.37.09a.82.82 0 0 1-.52-.2l-1.45-1.46a.69.69 0 0 1-.18-.71l1.93-6.27l-1.94-1.94m12.04-3.76l-10.1 5.7m-.66-16.47l-5.71 10.1"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13 35c4-.74 11.1-7.4 11.93-11.93C20.37 23.94 13.71 31 13 35Z"/>`),
		g.Group(children),
	)
}