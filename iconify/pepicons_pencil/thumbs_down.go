package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbsDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M14.1 10.081h2.5v-5.5h-2.5v5.5Zm2.5 1a1 1 0 0 0 1-1v-5.5a1 1 0 0 0-1-1h-2.5a1 1 0 0 0-1 1v5.5a1 1 0 0 0 1 1h2.5Zm-6.544 4.674a.55.55 0 0 1-.451-.632l.223-1.342a5.12 5.12 0 0 1 2.21-3.419l.61.914a4.021 4.021 0 0 0-1.736 2.685l-.224 1.342a.55.55 0 0 1-.632.452Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M8.384 15.42a.733.733 0 0 0 1.23-.34l1.065.266c-.345 1.381-2.065 1.858-3.071.851l-.047-.046a.55.55 0 1 1 .777-.777l.046.046Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M7.808 11.934a.55.55 0 0 1 .4.666l-.163.653c-.062.246-.089.5-.08.753l.015.451c.012.348.146.68.378.939l-.817.733a2.58 2.58 0 0 1-.658-1.634l-.016-.451a3.821 3.821 0 0 1 .112-1.058l.163-.652a.55.55 0 0 1 .666-.4Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M7.95 12.467a.55.55 0 0 1-.55.55H5.753a.55.55 0 1 1 0-1.1H7.4a.55.55 0 0 1 .55.55Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M3.86 10.748a2.117 2.117 0 0 0 1.893 1.17v1.098a3.215 3.215 0 0 1-2.876-1.777l-.362-.723a.55.55 0 1 1 .983-.491l.361.723Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M3.75 4.264a.55.55 0 0 1 .32.708L3.475 6.54a4.832 4.832 0 0 0 .04 3.524l-1.019.412a5.93 5.93 0 0 1-.049-4.325l.594-1.568a.55.55 0 0 1 .708-.319Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M6.702 3.659a2.122 2.122 0 0 0-2.385.93l-.293.477a.55.55 0 1 1-.936-.576l.293-.477a3.22 3.22 0 0 1 3.62-1.411l-.3 1.057Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M12.882 4.336a.55.55 0 0 1-.646.431L6.743 3.67a.55.55 0 1 1 .216-1.077L12.45 3.69a.55.55 0 0 1 .43.646Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M12.34 11.369a.55.55 0 0 1-.55-.55V4.23a.55.55 0 1 1 1.098 0v6.59a.55.55 0 0 1-.549.55Z" clip-rule="evenodd"/><path d="M14.35 6.081a1 1 0 1 1 2 0a1 1 0 0 1-2 0Z"/></g>`),
		g.Group(children),
	)
}