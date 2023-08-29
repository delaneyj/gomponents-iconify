package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToggleButtonStateB(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<g stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path fill="#d0cfce" d="M20.945 45.98h30.04c5.513 0 9.98-4.469 9.98-9.98c0-5.512-4.468-9.98-9.98-9.98h-30.04c-5.511 0-9.98 4.468-9.98 9.98c0 5.511 4.469 9.98 9.98 9.98z"/><circle cx="50.965" cy="36" r="10.001" fill="#b1cc33"/><path fill="#5c9e31" d="M43.962 43.072c3.905 3.905 10.238 3.905 14.143 0s3.906-10.238 0-14.144"/></g><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path d="M20.945 45.98h30.04c5.513 0 9.98-4.469 9.98-9.98v0c0-5.512-4.468-9.98-9.98-9.98h-30.04c-5.511 0-9.98 4.468-9.98 9.98v0c0 5.511 4.469 9.98 9.98 9.98z"/><circle cx="50.965" cy="36" r="10.001"/><path d="M43.962 43.072c3.905 3.905 10.238 3.905 14.143 0s3.906-10.238 0-14.144"/></g>`),
		g.Group(children),
	)
}