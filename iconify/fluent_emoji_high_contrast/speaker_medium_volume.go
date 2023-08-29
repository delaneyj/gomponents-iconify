package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeakerMediumVolume(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="m16.61 26.83l-8.612-3.829L3.803 23C2.807 23 2 22.153 2 21.108V10.892C2 9.847 2.807 9 3.803 9h4.21v.012l8.034-3.78c1.123-.529 2.989.023 2.989 1.264v6.474a3.055 3.055 0 0 1 0 6.107v6.227a1.693 1.693 0 0 1-2.425 1.527Zm.426-1.999v-17.8h-.03c-.067.001-.105.011-.113.013h-.002L9.998 10.29V21.7l7.038 3.13Z"/><path d="M24.11 11.07a7.3 7.3 0 0 0-.438-.442c-.637-.59-1.672-.137-1.672.732v.028c0 .271.118.528.317.713a4.986 4.986 0 0 1 1.604 3.667a4.98 4.98 0 0 1-1.604 3.668c-.205.19-.317.46-.317.74c0 .871 1.037 1.322 1.676.73a6.958 6.958 0 0 0 2.242-5.354a6.987 6.987 0 0 0-1.808-4.482Z"/></g>`),
		g.Group(children),
	)
}