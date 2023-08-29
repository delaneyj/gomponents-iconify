package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbsDownCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M17.1 13.081h2.5v-5.5h-2.5v5.5Zm2.5 1a1 1 0 0 0 1-1v-5.5a1 1 0 0 0-1-1h-2.5a1 1 0 0 0-1 1v5.5a1 1 0 0 0 1 1h2.5Zm-6.544 4.674a.55.55 0 0 1-.451-.632l.223-1.342a5.12 5.12 0 0 1 2.21-3.419l.61.914a4.021 4.021 0 0 0-1.736 2.686l-.224 1.341a.55.55 0 0 1-.632.452Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M11.384 18.42a.733.733 0 0 0 1.23-.34l1.065.266c-.345 1.381-2.065 1.858-3.071.851l-.047-.046a.55.55 0 1 1 .777-.777l.046.047Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M10.808 14.934a.55.55 0 0 1 .4.666l-.163.653c-.062.246-.089.5-.08.753l.015.451c.012.348.146.68.378.939l-.817.733a2.58 2.58 0 0 1-.658-1.634l-.016-.451a3.821 3.821 0 0 1 .112-1.058l.163-.652a.55.55 0 0 1 .666-.4Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M10.95 15.467a.55.55 0 0 1-.55.55H8.753a.55.55 0 1 1 0-1.1H10.4a.55.55 0 0 1 .55.55Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M6.86 13.748a2.117 2.117 0 0 0 1.893 1.17v1.098a3.215 3.215 0 0 1-2.876-1.777l-.362-.723a.55.55 0 1 1 .983-.491l.361.723Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M6.75 7.264a.55.55 0 0 1 .32.708L6.475 9.54a4.832 4.832 0 0 0 .04 3.524l-1.019.412a5.93 5.93 0 0 1-.049-4.325l.594-1.568a.55.55 0 0 1 .708-.319Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M9.702 6.659a2.122 2.122 0 0 0-2.385.93l-.293.477a.55.55 0 1 1-.936-.576l.293-.477a3.22 3.22 0 0 1 3.62-1.411l-.3 1.057Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M15.882 7.336a.55.55 0 0 1-.646.431L9.743 6.67a.55.55 0 1 1 .216-1.077L15.45 6.69a.55.55 0 0 1 .43.646Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M15.34 14.369a.55.55 0 0 1-.55-.55V7.23a.55.55 0 1 1 1.098 0v6.59a.55.55 0 0 1-.549.55Z" clip-rule="evenodd"/><path d="M17.35 9.081a1 1 0 1 1 2 0a1 1 0 0 1-2 0Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}