package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sonic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.273 17.007C8.446 13.247 21.387-.39 38.256 9.061c0 0-7.727 5.003-9.245 7.417c0 0 6.934-2.656 14.489 9.522c-3.346-.242-13.454 2.725-16.11 4.45"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.352 34.762c1.966.242 8.624 2.346 14.73 7.452c0 0 .54-9.623-4.318-13.298M6.898 27.692c-.828-2.38-1.449-8.281-.241-8.281s2.345 8.072 3.358 8.072c1.621 0 2.3-9.038 7.612-9.038c3.898 0 4.14 3.726 4.14 6.106c0 2.635-2.45 7.313-6.486 7.313c-1.633 0-3.967-2.948-5.692-2.948a5.85 5.85 0 0 0-.648.07m1.723-16.504c-1.248-1.661-3.076-2.8-3.594-2.8s-2.444 3.375-1.056 7.929m13.683-3.41c1.276-1.31 7.055-4.898 8.03-4.001s.13 9.142 0 9.728"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.113 18.962a14.519 14.519 0 0 0 1.104-7.037s-3.623 1.58-4.83 3.2c1.363.086 3.312 2.043 3.726 3.837Zm-7.41 12.267c.993-.52 3.77-1.314 3.77 1.463s-4.433 3.226-5.588 3.226c-2.38 0-6.83-1.087-8.642-5.572"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.534 33.658a4.033 4.033 0 0 0 4.14-1.127"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.466 32.428a4.24 4.24 0 0 1 1.208.103a3.654 3.654 0 0 1 .069 1.185m.855.804a5.862 5.862 0 0 0 3.894-2.414M17.604 21.86c.73 0 1.403 2 1.403 3.864s-.5 3.087-1.122 3.087s-1.431-.88-1.431-3.475s.42-3.476 1.15-3.476Zm-9.939-.87c-.112.543-.284 7.339 1.268 7.339c.398 0 .706-.362.943-.9M8.52 18.548c-.436-1.07-.811-1.578-1.579-1.578c-.923 0-1.665 1.803-1.665 3.269a5.836 5.836 0 0 0 .82 3.262"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.648 27.492c-.397.862-.076 1.811 1.957 2.613c1.487.586 2.444.33 2.668-.12c.437-.88-1.582-1.93-2.375-2.293s-1.964-.822-2.25-.2Z"/>`),
		g.Group(children),
	)
}