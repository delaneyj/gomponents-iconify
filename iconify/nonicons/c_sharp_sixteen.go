package nonicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CSharpSixteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M7 5.5a2.5 2.5 0 1 0 1.81 4.225a.75.75 0 0 1 1.086 1.035a4 4 0 1 1-.01-5.53a.75.75 0 0 1-1.082 1.04A2.49 2.49 0 0 0 7 5.5Z"/><path fill="currentColor" d="M6.586.102a.75.75 0 0 1 .756 0l4.715 2.75a.75.75 0 0 1-.756 1.296l-4.337-2.53L1.5 4.806v6.388l5.464 3.188l4.337-2.53a.75.75 0 1 1 .755 1.296l-4.714 2.75a.75.75 0 0 1-.756 0L.372 12.273A.75.75 0 0 1 0 11.625v-7.25a.75.75 0 0 1 .372-.648L6.586.102Z"/><path fill="currentColor" d="M12.18 5.25a.5.5 0 0 1 .5.5v4.5a.5.5 0 0 1-1 0v-4.5a.5.5 0 0 1 .5-.5Zm2.14 0a.5.5 0 0 1 .5.5v4.5a.5.5 0 0 1-1 0v-4.5a.5.5 0 0 1 .5-.5Z"/><path fill="currentColor" d="M10.5 6.93a.5.5 0 0 1 .5-.5h4.5a.5.5 0 0 1 0 1H11a.5.5 0 0 1-.5-.5Zm0 2.14a.5.5 0 0 1 .5-.5h4.5a.5.5 0 0 1 0 1H11a.5.5 0 0 1-.5-.5Z"/>`),
		g.Group(children),
	)
}