package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagEuropeanUnion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#1e50a0" d="M5 17h62v38H5z"/><path fill="#f1b31c" stroke="#f1b31c" stroke-linecap="round" stroke-linejoin="round" d="m35.105 26.794l.919-2.695l.792 2.654l-2.2-1.599l2.768-.067l-2.279 1.707zm0 21.107l.919-2.695l.792 2.653l-2.2-1.598l2.768-.067l-2.279 1.707zm-9.012-15.979l.918-2.695l.793 2.653l-2.201-1.598l2.768-.067l-2.278 1.707zm18.278 10.553l.919-2.695l.792 2.654l-2.2-1.599l2.768-.067l-2.279 1.707zM29.838 28.242l.919-2.695l.792 2.653l-2.201-1.598l2.768-.067l-2.278 1.707zM40.391 46.52l.919-2.694l.792 2.653l-2.201-1.599l2.769-.066l-2.279 1.706zm5.268-9.173l.919-2.694l.792 2.653l-2.201-1.599l2.768-.066l-2.278 1.706zm-21.107 0l.919-2.694l.792 2.653l-2.2-1.599l2.768-.066l-2.279 1.706zm15.729-9.131l.919-2.694l.792 2.653l-2.2-1.599l2.768-.066l-2.279 1.706zM29.728 46.495l.919-2.695l.792 2.654l-2.201-1.599l2.769-.067l-2.279 1.707zM44.13 31.886l.918-2.695l.793 2.654l-2.201-1.599l2.768-.067l-2.278 1.707zM25.851 42.439l.919-2.694l.792 2.653l-2.201-1.599l2.768-.066l-2.278 1.706z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}