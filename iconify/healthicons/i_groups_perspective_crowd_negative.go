package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IGroupsPerspectiveCrowdNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsIGroupsPerspectiveCrowdNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M0 0h48v48H0V0Zm18 16.5a4.5 4.5 0 0 1-4.5 4.5A4.499 4.499 0 0 1 9 16.5c0-2.486 2.014-4.5 4.5-4.5s4.5 2.014 4.5 4.5ZM13.5 23C10.33 23 4 24.787 4 28.333V36h19v-7.667C23 24.787 16.67 23 13.5 23ZM39 16.5a4.5 4.5 0 0 1-4.5 4.5a4.499 4.499 0 0 1-4.5-4.5c0-2.486 2.014-4.5 4.5-4.5s4.5 2.014 4.5 4.5ZM24 18a3 3 0 1 0 0-6a3 3 0 1 0 0 6Zm1 10.333C25 24.787 31.33 23 34.5 23s9.5 1.787 9.5 5.333V36H25v-7.667Zm2.962-5.835c.259-.11.522-.215.789-.313C27.228 21.396 25.27 21 24 21c-1.27 0-3.228.396-4.75 1.185c.266.098.53.202.787.313c1.172.5 2.354 1.176 3.273 2.08c.246.243.48.508.69.797c.21-.289.444-.554.69-.797c.919-.904 2.101-1.58 3.272-2.08Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsIGroupsPerspectiveCrowdNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}