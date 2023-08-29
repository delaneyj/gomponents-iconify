package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Godville(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m20.399 9.079l2.199 20.701l-5.103.756l1.29 2.61l-1.626 1.901l5.209 10.391"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m17.495 30.536l-1.31 2.043l-.568-1.745l-1.809 4.332l-.715-1.388l-1.198 2.019l-1.136-2.63l-1.009 1.704l-6.963-7.348m.518-9.372l2.933 6.985V14.453l5.552-5.341l11.775-5.383l-3.884-.794"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.542 16.85c.8-1.43 2.797-4.647 5.53-4.647s4.731 2.818 5.657 4.648m10.738-1.571v-3.077l11.776 1.832M27.509 8.678l-2.163 21.294l5.243.533M17.495 11.551l5.103-4.345l-6.645.785a2.516 2.516 0 0 0 0 2.215a3.832 3.832 0 0 0 1.542 1.345Zm2.744-4.067l-.163-2.16"/><ellipse cx="13.136" cy="17.647" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="4.383" ry="3.12"/><ellipse cx="35.173" cy="17.974" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.447" ry="4.383" transform="rotate(-84.769 35.173 17.974)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.867 17.647a7.766 7.766 0 0 0 4.383 1.344a7.4 7.4 0 0 0 4.269-1.344a7.4 7.4 0 0 0-4.269-1.344a7.766 7.766 0 0 0-4.383 1.344Z"/><circle cx="13.136" cy="17.647" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m42.109 12.407l-16.525-4.22l1.009-.589l-.546-1.556L40.029 9.67M26.512 2.645a3.1 3.1 0 0 0 .73 3.707m17.492 23.355a4.166 4.166 0 0 0-3.78-1.042a2.9 2.9 0 0 0-4.053-1.077c-.946-1.057-6.058-1.862-6.312 2.917c-1.412.505-1.626 2.72-.253 3.84a2.377 2.377 0 0 0-.61 2.748a4.054 4.054 0 0 0-2.305 6.253a1.36 1.36 0 0 0 .791 1.741m2.596-27.513c1.183 1.099 2.346 2.356 4.365 2.384a9.277 9.277 0 0 0 4.365-1.585"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.032 35.755c-.831-1.44-5.426-2.124-6.225 1.356M39.4 29.73c-.566-.432-3.168-.609-3.406 1.546m-3.66 11.271a3.16 3.16 0 0 1 4.29-2.86"/>`),
		g.Group(children),
	)
}