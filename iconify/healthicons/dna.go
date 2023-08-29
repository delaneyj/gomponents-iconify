package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dna(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M35 9.23V6h-2v3H15V6h-2v3.23a16.91 16.91 0 0 0 6.804 13.558a18.91 18.91 0 0 1 1.95-1.077A14.906 14.906 0 0 1 17.185 17h13.63A14.896 14.896 0 0 1 24 22.917l-.016.007c-.118.05-.236.1-.355.148A16.912 16.912 0 0 0 13.093 37H13v5h2v-3h18v3h2v-5h-.093a16.915 16.915 0 0 0-6.71-11.788a18.91 18.91 0 0 1-1.95 1.077A14.906 14.906 0 0 1 30.814 31h-13.63A14.896 14.896 0 0 1 24 25.083l.016-.007c.118-.05.236-.1.355-.148A16.91 16.91 0 0 0 35 9.23ZM31.838 33H16.162a14.876 14.876 0 0 0-1.056 4h17.79a14.876 14.876 0 0 0-1.057-4ZM16.162 15h15.677c.531-1.264.89-2.608 1.057-4h-17.79a14.876 14.876 0 0 0 1.056 4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}