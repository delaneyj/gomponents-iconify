package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SnowCamera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.48 43.5a5.788 5.788 0 0 1-3.442-1.138c-2.394-1.757-2.956-5.116-1.48-7.692l2.275-3.975a.616.616 0 0 0-.534-.922H15.87c-1.58 0-3.15-.54-4.273-1.653a5.794 5.794 0 0 1-.968-7.011l7.74-13.515c1.576-2.751 5.069-3.936 7.866-2.444a5.788 5.788 0 0 1 2.301 7.985l-2.37 4.14a.616.616 0 0 0 .534.922h5.431c1.58 0 3.15.541 4.271 1.654a5.795 5.795 0 0 1 .967 7.01l-7.862 13.727A5.785 5.785 0 0 1 24.48 43.5Z"/>`),
		g.Group(children),
	)
}