package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cajaarena(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Zm-5.78 17.618h7.78m-14.124 0h4.936m-27.812 0h14.567"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m42.5 29.919l-2.953-2.953h0l-.49.49l-5.041-5.041h0l-3.18 3.18l-6.623-6.623L8.957 34.228l-.49-.49L5.5 36.705"/><circle cx="24.49" cy="33.959" r=".75" fill="currentColor"/><circle cx="14.171" cy="33.959" r=".75" fill="currentColor"/><circle cx="26.449" cy="37.558" r=".75" fill="currentColor"/><circle cx="22.031" cy="38.922" r=".75" fill="currentColor"/><circle cx="9.73" cy="37.551" r=".75" fill="currentColor"/><circle cx="19.2" cy="32.718" r=".75" fill="currentColor"/><circle cx="21.159" cy="27.102" r=".75" fill="currentColor"/><circle cx="27.494" cy="27.428" r=".75" fill="currentColor"/><circle cx="24.189" cy="21.716" r=".75" fill="currentColor"/><circle cx="35.331" cy="27.559" r=".75" fill="currentColor"/><circle cx="35.396" cy="31.738" r=".75" fill="currentColor"/><circle cx="39.118" cy="34.677" r=".75" fill="currentColor"/><circle cx="37.616" cy="38.366" r=".75" fill="currentColor"/><circle cx="33.894" cy="40.032" r=".75" fill="currentColor"/><circle cx="30.824" cy="35.918" r=".75" fill="currentColor"/><circle cx="30.759" cy="30.824" r=".75" fill="currentColor"/><circle cx="29.126" cy="40.359" r=".75" fill="currentColor"/><circle cx="16.425" cy="37.616" r=".75" fill="currentColor"/><circle cx="22.031" cy="18.061" r=".75" fill="currentColor"/><circle cx="26.299" cy="18.061" r=".75" fill="currentColor"/><circle cx="24.356" cy="16.584" r=".75" fill="currentColor"/><circle cx="26.52" cy="13.425" r=".75" fill="currentColor"/><circle cx="24.577" cy="11.948" r=".75" fill="currentColor"/><circle cx="26.11" cy="10.138" r=".75" fill="currentColor"/><circle cx="26.389" cy="7.083" r=".75" fill="currentColor"/><circle cx="21.52" cy="8.228" r=".75" fill="currentColor"/><circle cx="23.576" cy="9.95" r=".75" fill="currentColor"/><circle cx="21.94" cy="11.572" r=".75" fill="currentColor"/><circle cx="22.073" cy="14.916" r=".75" fill="currentColor"/><circle cx="24" cy="14.166" r=".75" fill="currentColor"/><circle cx="24.326" cy="7.56" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}