package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WarNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsWarNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM7.136 22.005C6.986 23.1 7.903 24 9.007 24h22.984c1.105 0 2.023-.9 1.873-1.995a14.323 14.323 0 0 0-.75-3.005H43a1 1 0 1 0 0-2H32.194c-2.334-4.185-6.697-7-11.694-7c-6.803 0-12.43 5.218-13.364 12.005ZM7.187 26c-2.174 0-3.709 2.006-3.021 3.949l1.397 3.948C6.43 36.347 8.864 38 11.606 38h24.788c2.742 0 5.176-1.653 6.043-4.103l1.397-3.948c.688-1.943-.847-3.949-3.021-3.949H7.187Zm6.313 8.5a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm13-2.5a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0Zm8 2.5a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsWarNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}