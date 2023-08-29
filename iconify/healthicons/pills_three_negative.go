package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PillsThreeNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#healthiconsPills3Negative0)"><path d="M24.384 9.91a1 1 0 0 1 .914 1.08l-.602 7.187a1 1 0 0 1-1.994-.167l.603-7.187a1 1 0 0 1 1.08-.913Zm-6.627 26.847a1 1 0 0 1-1.414 0l-5.1-5.1a1 1 0 0 1 1.414-1.414l5.1 5.1a1 1 0 0 1 0 1.414Zm19.458-5.707a1 1 0 0 0-.762-1.848l-6.668 2.748a1 1 0 0 0 .762 1.849l6.668-2.748Z"/><path fill-rule="evenodd" d="M48 0H0v48h48V0ZM32.5 14.5a8.5 8.5 0 1 1-17 0a8.5 8.5 0 0 1 17 0Zm-9.5 19a8.5 8.5 0 1 1-17 0a8.5 8.5 0 0 1 17 0ZM33.5 40a8.5 8.5 0 1 0 0-17a8.5 8.5 0 0 0 0 17Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsPills3Negative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}