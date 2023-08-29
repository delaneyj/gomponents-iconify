package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BedLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M19 20v-1.5M5 20v-1.5"/><path d="M2 15c0-.932 0-1.398.152-1.765a2 2 0 0 1 1.083-1.083C3.602 12 4.068 12 5 12h14c.932 0 1.398 0 1.765.152a2 2 0 0 1 1.083 1.083C22 13.602 22 14.068 22 15c0 .932 0 1.398-.152 1.765a2 2 0 0 1-1.083 1.083C20.398 18 19.932 18 19 18H5c-.932 0-1.398 0-1.765-.152a2 2 0 0 1-1.083-1.083C2 16.398 2 15.932 2 15Zm19-3c0-3.771 0-5.657-1.172-6.828C18.657 4 16.771 4 13 4h-2C7.229 4 5.343 4 4.172 5.172C3 6.343 3 8.229 3 12"/><path d="M18.5 12v-1.5c0-1.886 0-2.828-.586-3.414C17.328 6.5 16.386 6.5 14.5 6.5h-5c-1.886 0-2.828 0-3.414.586C5.5 7.672 5.5 8.614 5.5 10.5V12M12 7v5"/></g>`),
		g.Group(children),
	)
}