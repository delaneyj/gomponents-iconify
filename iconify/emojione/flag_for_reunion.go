package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForReunion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#2a5f9e" d="M14 8v48l24-24z"/><path fill="#ffce31" d="M56 14C50.5 6.7 41.8 2 32 2c-6.8 0-13 2.2-18 6l24 24l18-18"/><path fill="#ed4c5c" d="M56 50c3.8-5 6-11.2 6-18c0-6.8-2.2-13-6-18L38 32l18 18"/><path fill="#699635" d="M38 32L14 56c5 3.8 11.2 6 18 6c9.8 0 18.5-4.7 24-12L38 32z"/><path fill="#2a5f9e" d="M6 17c-2.5 4.4-4 9.5-4 15s1.5 10.6 4 15V17"/><path fill="#fff" d="M6 17v30c1.1 2 2.5 3.8 4 5.4V11.6c-1.5 1.7-2.9 3.5-4 5.4"/><path fill="#ed4c5c" d="M10 11.6v40.8c1.2 1.3 2.6 2.5 4 3.6V8c-1.4 1.1-2.8 2.3-4 3.6"/><circle cx="38" cy="32" r="12" fill="#fff"/>`),
		g.Group(children),
	)
}