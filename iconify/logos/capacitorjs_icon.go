package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CapacitorjsIcon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#53B9FF" d="M39.863 54.115L.311 93.716l60.995 61.179L0 216.385l39.428 39.619l61.43-61.507l61.097 61.068l39.552-39.602z"/><path fill="#119EFF" d="m140.517 154.896l-39.658 39.601l61.097 61.069l39.552-39.602z"/><path fill-opacity=".2" d="m140.517 154.896l-39.658 39.601l15.267 15.182z"/><path fill="#53B9FF" d="M194.57 100.985L256 39.478L216.431 0l-61.412 61.384L93.917.311L54.365 39.913L216.01 201.761l39.552-39.602z"/><path fill="#119EFF" d="m115.36 100.987l39.659-39.602L93.917.313L54.365 39.914z"/><path fill-opacity=".2" d="m115.359 100.985l39.659-39.601l-15.271-15.186z"/>`),
		g.Group(children),
	)
}