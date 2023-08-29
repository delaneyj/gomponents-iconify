package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiplomaVerifiedLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M7 17.998c-2.175-.012-3.353-.108-4.121-.877C2 16.243 2 14.828 2 12V8c0-2.828 0-4.243.879-5.121C3.757 2 5.172 2 8 2h8c2.828 0 4.243 0 5.121.879C22 3.757 22 5.172 22 8v4c0 2.828 0 4.243-.879 5.121c-.73.73-1.829.854-3.801.875l-.82.002" opacity=".5"/><path stroke-linecap="round" d="M9 6h6M7 9.5h10" opacity=".5"/><path d="M10.89 13.945a1.71 1.71 0 0 1 2.22 0c.273.234.614.375.973.404a1.71 1.71 0 0 1 1.569 1.568c.028.36.17.7.403.974a1.71 1.71 0 0 1 0 2.218a1.71 1.71 0 0 0-.403.974a1.71 1.71 0 0 1-1.57 1.569a1.71 1.71 0 0 0-.973.403a1.71 1.71 0 0 1-2.218 0a1.71 1.71 0 0 0-.974-.404a1.71 1.71 0 0 1-1.568-1.568a1.71 1.71 0 0 0-.404-.974a1.71 1.71 0 0 1 0-2.218a1.71 1.71 0 0 0 .404-.974a1.71 1.71 0 0 1 1.568-1.568c.36-.029.7-.17.974-.404Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m10.5 18.2l.857.8l2.143-2"/></g>`),
		g.Group(children),
	)
}