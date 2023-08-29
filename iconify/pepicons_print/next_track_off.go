package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NextTrackOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><g opacity=".2"><path d="M11.64 10.747a1 1 0 0 1 0 1.505L6.159 17.05c-.647.566-1.659.106-1.659-.753V6.704a1 1 0 0 1 1.659-.753l5.48 4.796Z"/><path fill-rule="evenodd" d="m11.97 12.629l-5.482 4.796c-.97.849-2.488.16-2.488-1.129V6.704c0-1.289 1.518-1.978 2.488-1.13l5.481 4.797a1.5 1.5 0 0 1 0 2.258Zm-.33-.377a1 1 0 0 0 0-1.505L6.159 5.951c-.647-.566-1.659-.106-1.659.753v9.592a1 1 0 0 0 1.659.753l5.48-4.796Z" clip-rule="evenodd"/><path d="M18.64 10.747a1 1 0 0 1 0 1.505l-5.482 4.797c-.646.566-1.658.106-1.658-.753V6.704a1 1 0 0 1 1.659-.753l5.48 4.796Z"/><path fill-rule="evenodd" d="m18.97 12.629l-5.482 4.796c-.97.849-2.488.16-2.488-1.129V6.704c0-1.289 1.518-1.978 2.488-1.13l5.481 4.797a1.5 1.5 0 0 1 0 2.258Zm-.33-.377a1 1 0 0 0 0-1.505l-5.482-4.796c-.646-.566-1.658-.106-1.658.753v9.592a1 1 0 0 0 1.659.753l5.48-4.796Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M19 5.5a1 1 0 0 1 1 1v10a1 1 0 1 1-2 0v-10a1 1 0 0 1 1-1Z" clip-rule="evenodd"/></g><path fill-rule="evenodd" d="M3.5 14.796L8.981 10L3.5 5.204v9.592Zm6.14-4.043a1 1 0 0 0 0-1.506L4.159 4.451c-.647-.566-1.659-.106-1.659.753v9.592a1 1 0 0 0 1.659.753l5.48-4.796Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M10.5 14.796L15.981 10L10.5 5.204v9.592Zm6.14-4.043a1 1 0 0 0 0-1.506l-5.482-4.796c-.646-.566-1.658-.106-1.658.753v9.592a1 1 0 0 0 1.659.753l5.48-4.796Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M17 4.5a.5.5 0 0 1 .5.5v10a.5.5 0 0 1-1 0V5a.5.5 0 0 1 .5-.5Z" clip-rule="evenodd"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}