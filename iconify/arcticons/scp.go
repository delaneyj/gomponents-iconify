package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="25.401" r="11.21" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.728 19.688L24 23.632l2.272-3.944h-4.544Zm2.272 0v-7.266m-3.811 17.803l2.279-3.94l-4.552.004l2.272 3.936Zm-1.136-1.967L12.76 31.89m17.324-5.6l-4.552-.005l2.28 3.94l2.272-3.936Zm-1.136 1.967l6.292 3.633"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.584 28.251c.15-.93.25-1.878.25-2.85c0-7.62-4.788-14.105-11.512-16.659L29.26 4.766H18.741l-1.063 3.976c-6.724 2.554-11.511 9.04-11.511 16.659c0 .972.099 1.92.249 2.85L3.5 31.163l2.63 4.555l2.63 4.555l3.96-1.064c3.074 2.514 7 4.025 11.28 4.025s8.206-1.51 11.28-4.025l3.96 1.064l2.63-4.555l2.63-4.555l-2.916-2.912Z"/>`),
		g.Group(children),
	)
}