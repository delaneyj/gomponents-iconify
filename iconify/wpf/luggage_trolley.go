package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LuggageTrolley(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M15 3c-1.093 0-2 .907-2 2v1H8a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h17a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1h-5V5c0-1.093-.907-2-2-2h-3zM.812 4A1.001 1.001 0 0 0 1 6h1c.505 0 .66.108.781.25c.122.142.219.413.219.75v11c0 1.124.248 2.182.969 2.938C4.689 21.692 5.769 22 7 22h18a1 1 0 1 0 0-2H7c-.905 0-1.318-.212-1.563-.469C5.194 19.275 5 18.83 5 18V7c0-.69-.173-1.43-.688-2.031C3.798 4.368 2.95 4 2 4H1a1 1 0 0 0-.094 0a1.001 1.001 0 0 0-.094 0zM23 22a2 2 0 1 0 0 4a2 2 0 0 0 0-4zM7 22a2 2 0 1 0 0 4a2 2 0 0 0 0-4zm8-17h3v1h-3V5z"/>`),
		g.Group(children),
	)
}