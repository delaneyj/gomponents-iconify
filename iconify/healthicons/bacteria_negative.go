package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BacteriaNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsBacteriaNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM16 8.022a7.41 7.41 0 0 1 4.953 2.44l2.168-2.169l1.414 1.414l-2.44 2.44c.488.99.762 2.104.762 3.282a9.67 9.67 0 0 0 2.175 6.125l2.26-2.261l1.415 1.414l-2.261 2.261a9.674 9.674 0 0 0 6.125 2.175c.656 0 1.291.085 1.896.244l1.667-2.887l1.732 1l-1.537 2.662A7.425 7.425 0 0 1 40 32.572c0 1.087-.234 2.12-.654 3.05l2.618 1.512l-1 1.732l-2.679-1.547A7.413 7.413 0 0 1 32.571 40c-.19 0-.381-.002-.571-.007V43h-2v-3.133a24.411 24.411 0 0 1-8.606-2.55l-1.528 2.647l-1.732-1l1.518-2.63a24.693 24.693 0 0 1-5.312-4.432l-2.633 2.634l-1.414-1.415l2.762-2.762a24.509 24.509 0 0 1-3.801-7.16L6.259 24l-.518-1.932l2.96-.793A24.631 24.631 0 0 1 8 15.43c0-.681.092-1.34.263-1.967L5.5 11.866l1-1.732l2.556 1.476A7.437 7.437 0 0 1 14 8.137V5h2v3.022ZM14 16a1 1 0 1 1 2 0a1 1 0 0 1-2 0Zm1-3a3 3 0 1 0 0 6a3 3 0 0 0 0-6Zm8 17a1 1 0 1 1 2 0a1 1 0 0 1-2 0Zm1-3a3 3 0 1 0 0 6a3 3 0 0 0 0-6Zm-5-3a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm13 10a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsBacteriaNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}