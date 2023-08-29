package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gitlab(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m256 491.816l94.272-290.14H161.734L256 491.816zM29.606 201.676l-28.65 88.17a19.517 19.517 0 0 0 7.093 21.822L256 491.816L29.606 201.676z"/><path fill="currentColor" d="M29.606 201.676H161.73L104.946 26.927c-2.92-8.991-15.64-8.991-18.56 0l-56.78 174.75zm452.787 0l28.651 88.17a19.517 19.517 0 0 1-7.093 21.822l-247.95 180.148l226.392-290.14z"/><path fill="currentColor" d="M482.394 201.676H350.271l56.783-174.749c2.922-8.991 15.642-8.991 18.562 0l56.778 174.75z"/>`),
		g.Group(children),
	)
}