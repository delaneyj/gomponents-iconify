package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Useralt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024 958q0 12-13.5 22T969 996.5t-57.5 12t-75.5 8.5t-80 4.5t-87.5 2.5t-81 1h-151l-81-1l-87.5-2.5l-80-4.5l-75.5-8.5l-57.5-12L13.5 980L0 958q2-88 110-155.5T384 713v-33q-52-23-90-65t-59.5-98.5t-32-121T192 257q0-65 25-115t69-80.5t101-46T512 0t125 15.5t101 46t69 80.5t25 115q0 349-192 426v30q166 22 274 89.5T1024 958zM768 274q0-104-71.5-157T512 64t-184.5 53T256 274q0 68 10 125t32 106.5t60 82.5t90 46v138q-60 5-117 20.5t-98 34t-73 38t-48 34.5t-16 22q0 14 38 23t114 12t127 4t137 1t137-1t127-4t114-12t38-23q0-7-16-22t-48-34.5t-72.5-38t-98-34T576 772V634q192-41 192-360z"/>`),
		g.Group(children),
	)
}