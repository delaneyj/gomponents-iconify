package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FunWithKanji(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.805 20.956v-3.424h15.124v3.424m-7.61-3.424v-2.569m-8.085 11.7H43.5M9.637 16.771h12.651m-8.847 1.712v-3.52m5.042 3.52v-3.52m-7.324 5.422h9.512v2.949h-9.512zM4.785 32.276c.666-.952 1.522-2.283 2.093-3.9c.38-.952.57-1.808.761-2.569M4.5 20.1c.476.19.951.57 1.522.951c.666.476 1.141.951 1.522 1.332M5.737 15.82c.475.19.95.57 1.522.95c.666.476 1.141.952 1.522 1.332m1.807 7.896h10.749m1.331 2.948H9.256"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.01 20.576c.095 1.712 0 3.044 0 4.09c-.095 2.663-.38 3.33-.57 3.71c-.667 1.522-1.808 2.378-2.379 2.758c-1.712 1.237-3.424 1.522-4.376 1.617m13.888.286c-.76-.19-1.997-.476-3.234-1.332c-1.427-.856-2.283-1.902-2.854-2.568m13.888-7.8h9.893a12.425 12.425 0 0 1-2.283 2.092a15.31 15.31 0 0 1-2.663 1.713v6.563c0 .095 0 .476-.286.76c-.38.477-1.046.477-1.141.477H31.8"/>`),
		g.Group(children),
	)
}