package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wheel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<g fill="#3f3f3f"><path d="M36 8.071C20.575 8.071 8.071 20.575 8.071 36S20.575 63.929 36 63.929S63.929 51.425 63.929 36S51.425 8.071 36 8.071Zm0 48.845c-11.551 0-20.916-9.365-20.916-20.916S24.45 15.084 36 15.084S56.916 24.45 56.916 36S47.55 56.916 36 56.916Z"/><path d="M36 29.244a6.756 6.756 0 1 0 0 13.512a6.756 6.756 0 0 0 0-13.512ZM36 39a3 3 0 1 1 0-6a3 3 0 0 1 0 6Z"/></g><g fill="none" stroke="#000" stroke-miterlimit="10" stroke-width="2"><circle cx="36" cy="36" r="27.929"/><circle cx="36" cy="36" r="21.298"/><path stroke-linecap="round" d="m21.032 50.968l10.191-10.191m10.255 15.669l-3.729-13.92m18.697-1.048l-13.92-3.729m8.442-16.717L40.777 31.223M30.521 15.554l3.731 13.921m-18.698 1.046l13.919 3.73"/><circle cx="36" cy="36" r="6.756"/><circle cx="36" cy="36" r="3"/></g>`),
		g.Group(children),
	)
}