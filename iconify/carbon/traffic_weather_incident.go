package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrafficWeatherIncident(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M2 24h2v6H2zM28 2h2v28h-2zM15 2h2v4h-2zm0 8h2v4h-2zm0 8h2v4h-2zm0 8h2v4h-2zM6 12a3.898 3.898 0 0 1-4-3.777a3.902 3.902 0 0 1 .653-2.064L5.17 2.414a1.038 1.038 0 0 1 1.66 0L9.314 6.11A3.97 3.97 0 0 1 10 8.223A3.898 3.898 0 0 1 6 12zm0-7.237L4.344 7.226A1.89 1.89 0 0 0 4 8.223a1.9 1.9 0 0 0 2 1.778a1.9 1.9 0 0 0 2-1.778a1.98 1.98 0 0 0-.375-1.047zm5 7l-1.656 2.462a1.89 1.89 0 0 0-.344.998A1.9 1.9 0 0 0 11 17a1.9 1.9 0 0 0 2-1.777a1.98 1.98 0 0 0-.375-1.047zm-5 4l-1.656 2.462a1.89 1.89 0 0 0-.344.998A1.9 1.9 0 0 0 6 21a1.9 1.9 0 0 0 2-1.777a1.98 1.98 0 0 0-.375-1.047z"/>`),
		g.Group(children),
	)
}