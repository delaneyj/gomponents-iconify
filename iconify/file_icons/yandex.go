package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yandex(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M200.01 319.442V512H256V0h-83.63C90.186 0 21.09 55.511 21.09 163.677c0 77.168 30.552 119 76.374 142.073L0 512h64.73l88.731-192.558h46.55zm-.175-44.918h-29.81c-48.733 0-88.746-26.684-88.746-109.62c0-85.808 43.638-116.441 88.745-116.441h29.811v226.061z"/>`),
		g.Group(children),
	)
}