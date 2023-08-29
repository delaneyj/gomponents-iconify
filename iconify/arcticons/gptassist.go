package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gptassist(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.53 31.569V7.298c0-.993-.806-1.798-1.799-1.798H8.27c-.993 0-1.798.805-1.798 1.798v24.271c0 .993.805 1.798 1.798 1.798h3.933l3.932 9.133l3.933-9.133h19.664c.993 0 1.798-.805 1.798-1.798Z"/><circle cx="17.258" cy="15.462" r="2.247" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="30.742" cy="15.462" r="2.247" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.99 22.464c0 1.868-1.714 3.594-4.495 4.528c-2.782.934-6.208.934-8.99 0c-2.781-.934-4.494-2.66-4.494-4.528h17.978Z"/>`),
		g.Group(children),
	)
}