package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Axe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M7.038 8.5c-.671.672-2.108.652-4.198-1.438C.752 4.974.793 3.474 1.402 2.865c.61-.61.883.455 2.363-.471C5.244 1.466 5.564.487 6.081 1.003L8.897 3.82c.518.517-.797 1.173-1.39 2.316c-.549 1.065.202 1.694-.469 2.364zm9.259 6.457c-.142.143-.461.053-.715-.201l-6.961-6.96c-.252-.253-.344-.573-.201-.716l.577-.576c.141-.143.462-.051.716.201l6.959 6.961c.254.253.344.572.201.715l-.576.576z"/><path d="M3.008.137h.972v1.702h-.972z"/></g>`),
		g.Group(children),
	)
}