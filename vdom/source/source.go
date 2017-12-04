package source

import (
	"encoding/json"
	"strings"

	"github.com/matthewmueller/joy/dom/window"
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/vdom"
)

// Source struct
// js:"source,omit"
type Source struct {
	vdom.Child
	vdom.Node

	attrs    map[string]interface{}
	children []vdom.Child
}

// Props struct
// js:"props,omit"
type Props struct {
	attrs map[string]interface{}
}

// New source element
//
// The source element allows authors to specify multiple alternative media resources for media elements. It does not represent anything on its own.
func New(props *Props, children ...vdom.Child) *Source {
	macro.Rewrite("$1('source', $2 ? $2.JSON() : {}, $3)", vdom.Pragma(), props, children)
	if props == nil {
		props = &Props{attrs: map[string]interface{}{}}
	}
	return &Source{
		attrs:    props.attrs,
		children: children,
	}
}

// Render fn
func (s *Source) Render() vdom.Node {
	return s
}

func (s *Source) String() string {
	// props
	var props []string
	for key, val := range s.attrs {
		bytes, e := json.Marshal(val)
		// TODO: skip over errors?
		if e != nil {
			continue
		}
		props = append(props, key+"="+string(bytes))
	}

	// children
	var children []string
	for _, child := range s.children {
		children = append(children, child.Render().String())
	}

	if len(props) > 0 {
		return "<source " + strings.Join(props, " ") + ">" + strings.Join(children, "") + "</source>"
	}

	return "<source>" + strings.Join(children, "") + "</source>"
}

// Src fn
func Src(src string) *Props {
	macro.Rewrite("$1().Set('src', $2)", macro.Runtime("Map", "Set", "JSON"), src)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Src(src)
}

// Src fn
func (p *Props) Src(src string) *Props {
	macro.Rewrite("$_.Set('src', $1)", src)
	p.attrs["src"] = src
	return p
}

// Type fn
func Type(kind string) *Props {
	macro.Rewrite("$1().Set('type', $2)", macro.Runtime("Map", "Set", "JSON"), kind)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Type(kind)
}

// Type fn
func (p *Props) Type(kind string) *Props {
	macro.Rewrite("$_.Set('type', $1)", kind)
	p.attrs["type"] = kind
	return p
}

// Ondrop fn
func Ondrop(ondrop func(e window.DragEvent)) *Props {
	macro.Rewrite("$1().Set('ondrop', $2)", macro.Runtime("Map", "Set", "JSON"), ondrop)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Ondrop(ondrop)
}

// Ondrop fn
func (p *Props) Ondrop(ondrop func(e window.DragEvent)) *Props {
	macro.Rewrite("$_.Set('ondrop', $1)", ondrop)
	p.attrs["ondrop"] = ondrop
	return p
}

// Onresize fn
func Onresize(onresize func(e window.UIEvent)) *Props {
	macro.Rewrite("$1().Set('onresize', $2)", macro.Runtime("Map", "Set", "JSON"), onresize)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onresize(onresize)
}

// Onresize fn
func (p *Props) Onresize(onresize func(e window.UIEvent)) *Props {
	macro.Rewrite("$_.Set('onresize', $1)", onresize)
	p.attrs["onresize"] = onresize
	return p
}

// Onseeking fn
func Onseeking(onseeking func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('onseeking', $2)", macro.Runtime("Map", "Set", "JSON"), onseeking)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onseeking(onseeking)
}

// Onseeking fn
func (p *Props) Onseeking(onseeking func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('onseeking', $1)", onseeking)
	p.attrs["onseeking"] = onseeking
	return p
}

// Onmousedown fn
func Onmousedown(onmousedown func(e window.MouseEvent)) *Props {
	macro.Rewrite("$1().Set('onmousedown', $2)", macro.Runtime("Map", "Set", "JSON"), onmousedown)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onmousedown(onmousedown)
}

// Onmousedown fn
func (p *Props) Onmousedown(onmousedown func(e window.MouseEvent)) *Props {
	macro.Rewrite("$_.Set('onmousedown', $1)", onmousedown)
	p.attrs["onmousedown"] = onmousedown
	return p
}

// Onmousemove fn
func Onmousemove(onmousemove func(e window.MouseEvent)) *Props {
	macro.Rewrite("$1().Set('onmousemove', $2)", macro.Runtime("Map", "Set", "JSON"), onmousemove)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onmousemove(onmousemove)
}

// Onmousemove fn
func (p *Props) Onmousemove(onmousemove func(e window.MouseEvent)) *Props {
	macro.Rewrite("$_.Set('onmousemove', $1)", onmousemove)
	p.attrs["onmousemove"] = onmousemove
	return p
}

// Onreset fn
func Onreset(onreset func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('onreset', $2)", macro.Runtime("Map", "Set", "JSON"), onreset)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onreset(onreset)
}

// Onreset fn
func (p *Props) Onreset(onreset func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('onreset', $1)", onreset)
	p.attrs["onreset"] = onreset
	return p
}

// Onreadystatechange fn
func Onreadystatechange(onreadystatechange func(e window.ProgressEvent)) *Props {
	macro.Rewrite("$1().Set('onreadystatechange', $2)", macro.Runtime("Map", "Set", "JSON"), onreadystatechange)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onreadystatechange(onreadystatechange)
}

// Onreadystatechange fn
func (p *Props) Onreadystatechange(onreadystatechange func(e window.ProgressEvent)) *Props {
	macro.Rewrite("$_.Set('onreadystatechange', $1)", onreadystatechange)
	p.attrs["onreadystatechange"] = onreadystatechange
	return p
}

// Onseeked fn
func Onseeked(onseeked func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('onseeked', $2)", macro.Runtime("Map", "Set", "JSON"), onseeked)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onseeked(onseeked)
}

// Onseeked fn
func (p *Props) Onseeked(onseeked func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('onseeked', $1)", onseeked)
	p.attrs["onseeked"] = onseeked
	return p
}

// Oncanplay fn
func Oncanplay(oncanplay func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('oncanplay', $2)", macro.Runtime("Map", "Set", "JSON"), oncanplay)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Oncanplay(oncanplay)
}

// Oncanplay fn
func (p *Props) Oncanplay(oncanplay func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('oncanplay', $1)", oncanplay)
	p.attrs["oncanplay"] = oncanplay
	return p
}

// Oncanplaythrough fn
func Oncanplaythrough(oncanplaythrough func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('oncanplaythrough', $2)", macro.Runtime("Map", "Set", "JSON"), oncanplaythrough)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Oncanplaythrough(oncanplaythrough)
}

// Oncanplaythrough fn
func (p *Props) Oncanplaythrough(oncanplaythrough func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('oncanplaythrough', $1)", oncanplaythrough)
	p.attrs["oncanplaythrough"] = oncanplaythrough
	return p
}

// Ondragend fn
func Ondragend(ondragend func(e window.DragEvent)) *Props {
	macro.Rewrite("$1().Set('ondragend', $2)", macro.Runtime("Map", "Set", "JSON"), ondragend)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Ondragend(ondragend)
}

// Ondragend fn
func (p *Props) Ondragend(ondragend func(e window.DragEvent)) *Props {
	macro.Rewrite("$_.Set('ondragend', $1)", ondragend)
	p.attrs["ondragend"] = ondragend
	return p
}

// Oninvalid fn
func Oninvalid(oninvalid func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('oninvalid', $2)", macro.Runtime("Map", "Set", "JSON"), oninvalid)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Oninvalid(oninvalid)
}

// Oninvalid fn
func (p *Props) Oninvalid(oninvalid func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('oninvalid', $1)", oninvalid)
	p.attrs["oninvalid"] = oninvalid
	return p
}

// Onplay fn
func Onplay(onplay func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('onplay', $2)", macro.Runtime("Map", "Set", "JSON"), onplay)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onplay(onplay)
}

// Onplay fn
func (p *Props) Onplay(onplay func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('onplay', $1)", onplay)
	p.attrs["onplay"] = onplay
	return p
}

// Onselect fn
func Onselect(onselect func(e window.UIEvent)) *Props {
	macro.Rewrite("$1().Set('onselect', $2)", macro.Runtime("Map", "Set", "JSON"), onselect)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onselect(onselect)
}

// Onselect fn
func (p *Props) Onselect(onselect func(e window.UIEvent)) *Props {
	macro.Rewrite("$_.Set('onselect', $1)", onselect)
	p.attrs["onselect"] = onselect
	return p
}

// Oncontextmenu fn
func Oncontextmenu(oncontextmenu func(e window.PointerEvent)) *Props {
	macro.Rewrite("$1().Set('oncontextmenu', $2)", macro.Runtime("Map", "Set", "JSON"), oncontextmenu)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Oncontextmenu(oncontextmenu)
}

// Oncontextmenu fn
func (p *Props) Oncontextmenu(oncontextmenu func(e window.PointerEvent)) *Props {
	macro.Rewrite("$_.Set('oncontextmenu', $1)", oncontextmenu)
	p.attrs["oncontextmenu"] = oncontextmenu
	return p
}

// Ondrag fn
func Ondrag(ondrag func(e window.DragEvent)) *Props {
	macro.Rewrite("$1().Set('ondrag', $2)", macro.Runtime("Map", "Set", "JSON"), ondrag)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Ondrag(ondrag)
}

// Ondrag fn
func (p *Props) Ondrag(ondrag func(e window.DragEvent)) *Props {
	macro.Rewrite("$_.Set('ondrag', $1)", ondrag)
	p.attrs["ondrag"] = ondrag
	return p
}

// Ondurationchange fn
func Ondurationchange(ondurationchange func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('ondurationchange', $2)", macro.Runtime("Map", "Set", "JSON"), ondurationchange)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Ondurationchange(ondurationchange)
}

// Ondurationchange fn
func (p *Props) Ondurationchange(ondurationchange func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('ondurationchange', $1)", ondurationchange)
	p.attrs["ondurationchange"] = ondurationchange
	return p
}

// Ondragenter fn
func Ondragenter(ondragenter func(e window.DragEvent)) *Props {
	macro.Rewrite("$1().Set('ondragenter', $2)", macro.Runtime("Map", "Set", "JSON"), ondragenter)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Ondragenter(ondragenter)
}

// Ondragenter fn
func (p *Props) Ondragenter(ondragenter func(e window.DragEvent)) *Props {
	macro.Rewrite("$_.Set('ondragenter', $1)", ondragenter)
	p.attrs["ondragenter"] = ondragenter
	return p
}

// Onloadeddata fn
func Onloadeddata(onloadeddata func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('onloadeddata', $2)", macro.Runtime("Map", "Set", "JSON"), onloadeddata)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onloadeddata(onloadeddata)
}

// Onloadeddata fn
func (p *Props) Onloadeddata(onloadeddata func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('onloadeddata', $1)", onloadeddata)
	p.attrs["onloadeddata"] = onloadeddata
	return p
}

// Onloadedmetadata fn
func Onloadedmetadata(onloadedmetadata func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('onloadedmetadata', $2)", macro.Runtime("Map", "Set", "JSON"), onloadedmetadata)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onloadedmetadata(onloadedmetadata)
}

// Onloadedmetadata fn
func (p *Props) Onloadedmetadata(onloadedmetadata func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('onloadedmetadata', $1)", onloadedmetadata)
	p.attrs["onloadedmetadata"] = onloadedmetadata
	return p
}

// Onmouseup fn
func Onmouseup(onmouseup func(e window.MouseEvent)) *Props {
	macro.Rewrite("$1().Set('onmouseup', $2)", macro.Runtime("Map", "Set", "JSON"), onmouseup)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onmouseup(onmouseup)
}

// Onmouseup fn
func (p *Props) Onmouseup(onmouseup func(e window.MouseEvent)) *Props {
	macro.Rewrite("$_.Set('onmouseup', $1)", onmouseup)
	p.attrs["onmouseup"] = onmouseup
	return p
}

// Onprogress fn
func Onprogress(onprogress func(e window.ProgressEvent)) *Props {
	macro.Rewrite("$1().Set('onprogress', $2)", macro.Runtime("Map", "Set", "JSON"), onprogress)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onprogress(onprogress)
}

// Onprogress fn
func (p *Props) Onprogress(onprogress func(e window.ProgressEvent)) *Props {
	macro.Rewrite("$_.Set('onprogress', $1)", onprogress)
	p.attrs["onprogress"] = onprogress
	return p
}

// Onshow fn
func Onshow(onshow func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('onshow', $2)", macro.Runtime("Map", "Set", "JSON"), onshow)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onshow(onshow)
}

// Onshow fn
func (p *Props) Onshow(onshow func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('onshow', $1)", onshow)
	p.attrs["onshow"] = onshow
	return p
}

// Onsubmit fn
func Onsubmit(onsubmit func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('onsubmit', $2)", macro.Runtime("Map", "Set", "JSON"), onsubmit)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onsubmit(onsubmit)
}

// Onsubmit fn
func (p *Props) Onsubmit(onsubmit func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('onsubmit', $1)", onsubmit)
	p.attrs["onsubmit"] = onsubmit
	return p
}

// Onblur fn
func Onblur(onblur func(e window.FocusEvent)) *Props {
	macro.Rewrite("$1().Set('onblur', $2)", macro.Runtime("Map", "Set", "JSON"), onblur)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onblur(onblur)
}

// Onblur fn
func (p *Props) Onblur(onblur func(e window.FocusEvent)) *Props {
	macro.Rewrite("$_.Set('onblur', $1)", onblur)
	p.attrs["onblur"] = onblur
	return p
}

// Oninput fn
func Oninput(oninput func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('oninput', $2)", macro.Runtime("Map", "Set", "JSON"), oninput)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Oninput(oninput)
}

// Oninput fn
func (p *Props) Oninput(oninput func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('oninput', $1)", oninput)
	p.attrs["oninput"] = oninput
	return p
}

// Onload fn
func Onload(onload func(e window.ProgressEvent)) *Props {
	macro.Rewrite("$1().Set('onload', $2)", macro.Runtime("Map", "Set", "JSON"), onload)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onload(onload)
}

// Onload fn
func (p *Props) Onload(onload func(e window.ProgressEvent)) *Props {
	macro.Rewrite("$_.Set('onload', $1)", onload)
	p.attrs["onload"] = onload
	return p
}

// Ontimeupdate fn
func Ontimeupdate(ontimeupdate func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('ontimeupdate', $2)", macro.Runtime("Map", "Set", "JSON"), ontimeupdate)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Ontimeupdate(ontimeupdate)
}

// Ontimeupdate fn
func (p *Props) Ontimeupdate(ontimeupdate func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('ontimeupdate', $1)", ontimeupdate)
	p.attrs["ontimeupdate"] = ontimeupdate
	return p
}

// Onloadstart fn
func Onloadstart(onloadstart func(e window.ProgressEvent)) *Props {
	macro.Rewrite("$1().Set('onloadstart', $2)", macro.Runtime("Map", "Set", "JSON"), onloadstart)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onloadstart(onloadstart)
}

// Onloadstart fn
func (p *Props) Onloadstart(onloadstart func(e window.ProgressEvent)) *Props {
	macro.Rewrite("$_.Set('onloadstart', $1)", onloadstart)
	p.attrs["onloadstart"] = onloadstart
	return p
}

// Onwaiting fn
func Onwaiting(onwaiting func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('onwaiting', $2)", macro.Runtime("Map", "Set", "JSON"), onwaiting)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onwaiting(onwaiting)
}

// Onwaiting fn
func (p *Props) Onwaiting(onwaiting func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('onwaiting', $1)", onwaiting)
	p.attrs["onwaiting"] = onwaiting
	return p
}

// Onended fn
func Onended(onended func(e window.MediaStreamErrorEvent)) *Props {
	macro.Rewrite("$1().Set('onended', $2)", macro.Runtime("Map", "Set", "JSON"), onended)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onended(onended)
}

// Onended fn
func (p *Props) Onended(onended func(e window.MediaStreamErrorEvent)) *Props {
	macro.Rewrite("$_.Set('onended', $1)", onended)
	p.attrs["onended"] = onended
	return p
}

// Onmouseout fn
func Onmouseout(onmouseout func(e window.MouseEvent)) *Props {
	macro.Rewrite("$1().Set('onmouseout', $2)", macro.Runtime("Map", "Set", "JSON"), onmouseout)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onmouseout(onmouseout)
}

// Onmouseout fn
func (p *Props) Onmouseout(onmouseout func(e window.MouseEvent)) *Props {
	macro.Rewrite("$_.Set('onmouseout', $1)", onmouseout)
	p.attrs["onmouseout"] = onmouseout
	return p
}

// Onclick fn
func Onclick(onclick func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('onclick', $2)", macro.Runtime("Map", "Set", "JSON"), onclick)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onclick(onclick)
}

// Onclick fn
func (p *Props) Onclick(onclick func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('onclick', $1)", onclick)
	p.attrs["onclick"] = onclick
	return p
}

// Onplaying fn
func Onplaying(onplaying func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('onplaying', $2)", macro.Runtime("Map", "Set", "JSON"), onplaying)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onplaying(onplaying)
}

// Onplaying fn
func (p *Props) Onplaying(onplaying func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('onplaying', $1)", onplaying)
	p.attrs["onplaying"] = onplaying
	return p
}

// Onsuspend fn
func Onsuspend(onsuspend func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('onsuspend', $2)", macro.Runtime("Map", "Set", "JSON"), onsuspend)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onsuspend(onsuspend)
}

// Onsuspend fn
func (p *Props) Onsuspend(onsuspend func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('onsuspend', $1)", onsuspend)
	p.attrs["onsuspend"] = onsuspend
	return p
}

// Onabort fn
func Onabort(onabort func(e window.ProgressEvent)) *Props {
	macro.Rewrite("$1().Set('onabort', $2)", macro.Runtime("Map", "Set", "JSON"), onabort)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onabort(onabort)
}

// Onabort fn
func (p *Props) Onabort(onabort func(e window.ProgressEvent)) *Props {
	macro.Rewrite("$_.Set('onabort', $1)", onabort)
	p.attrs["onabort"] = onabort
	return p
}

// Onerror fn
func Onerror(onerror func(e window.ProgressEvent)) *Props {
	macro.Rewrite("$1().Set('onerror', $2)", macro.Runtime("Map", "Set", "JSON"), onerror)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onerror(onerror)
}

// Onerror fn
func (p *Props) Onerror(onerror func(e window.ProgressEvent)) *Props {
	macro.Rewrite("$_.Set('onerror', $1)", onerror)
	p.attrs["onerror"] = onerror
	return p
}

// Onmousewheel fn
func Onmousewheel(onmousewheel func(e window.WheelEvent)) *Props {
	macro.Rewrite("$1().Set('onmousewheel', $2)", macro.Runtime("Map", "Set", "JSON"), onmousewheel)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onmousewheel(onmousewheel)
}

// Onmousewheel fn
func (p *Props) Onmousewheel(onmousewheel func(e window.WheelEvent)) *Props {
	macro.Rewrite("$_.Set('onmousewheel', $1)", onmousewheel)
	p.attrs["onmousewheel"] = onmousewheel
	return p
}

// Onkeyup fn
func Onkeyup(onkeyup func(e window.KeyboardEvent)) *Props {
	macro.Rewrite("$1().Set('onkeyup', $2)", macro.Runtime("Map", "Set", "JSON"), onkeyup)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onkeyup(onkeyup)
}

// Onkeyup fn
func (p *Props) Onkeyup(onkeyup func(e window.KeyboardEvent)) *Props {
	macro.Rewrite("$_.Set('onkeyup', $1)", onkeyup)
	p.attrs["onkeyup"] = onkeyup
	return p
}

// Ondragleave fn
func Ondragleave(ondragleave func(e window.DragEvent)) *Props {
	macro.Rewrite("$1().Set('ondragleave', $2)", macro.Runtime("Map", "Set", "JSON"), ondragleave)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Ondragleave(ondragleave)
}

// Ondragleave fn
func (p *Props) Ondragleave(ondragleave func(e window.DragEvent)) *Props {
	macro.Rewrite("$_.Set('ondragleave', $1)", ondragleave)
	p.attrs["ondragleave"] = ondragleave
	return p
}

// Onvolumechange fn
func Onvolumechange(onvolumechange func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('onvolumechange', $2)", macro.Runtime("Map", "Set", "JSON"), onvolumechange)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onvolumechange(onvolumechange)
}

// Onvolumechange fn
func (p *Props) Onvolumechange(onvolumechange func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('onvolumechange', $1)", onvolumechange)
	p.attrs["onvolumechange"] = onvolumechange
	return p
}

// Onmouseover fn
func Onmouseover(onmouseover func(e window.MouseEvent)) *Props {
	macro.Rewrite("$1().Set('onmouseover', $2)", macro.Runtime("Map", "Set", "JSON"), onmouseover)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onmouseover(onmouseover)
}

// Onmouseover fn
func (p *Props) Onmouseover(onmouseover func(e window.MouseEvent)) *Props {
	macro.Rewrite("$_.Set('onmouseover', $1)", onmouseover)
	p.attrs["onmouseover"] = onmouseover
	return p
}

// Onpause fn
func Onpause(onpause func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('onpause', $2)", macro.Runtime("Map", "Set", "JSON"), onpause)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onpause(onpause)
}

// Onpause fn
func (p *Props) Onpause(onpause func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('onpause', $1)", onpause)
	p.attrs["onpause"] = onpause
	return p
}

// Onemptied fn
func Onemptied(onemptied func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('onemptied', $2)", macro.Runtime("Map", "Set", "JSON"), onemptied)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onemptied(onemptied)
}

// Onemptied fn
func (p *Props) Onemptied(onemptied func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('onemptied', $1)", onemptied)
	p.attrs["onemptied"] = onemptied
	return p
}

// Onfocus fn
func Onfocus(onfocus func(e window.FocusEvent)) *Props {
	macro.Rewrite("$1().Set('onfocus', $2)", macro.Runtime("Map", "Set", "JSON"), onfocus)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onfocus(onfocus)
}

// Onfocus fn
func (p *Props) Onfocus(onfocus func(e window.FocusEvent)) *Props {
	macro.Rewrite("$_.Set('onfocus', $1)", onfocus)
	p.attrs["onfocus"] = onfocus
	return p
}

// Onkeypress fn
func Onkeypress(onkeypress func(e window.KeyboardEvent)) *Props {
	macro.Rewrite("$1().Set('onkeypress', $2)", macro.Runtime("Map", "Set", "JSON"), onkeypress)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onkeypress(onkeypress)
}

// Onkeypress fn
func (p *Props) Onkeypress(onkeypress func(e window.KeyboardEvent)) *Props {
	macro.Rewrite("$_.Set('onkeypress', $1)", onkeypress)
	p.attrs["onkeypress"] = onkeypress
	return p
}

// Onchange fn
func Onchange(onchange func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('onchange', $2)", macro.Runtime("Map", "Set", "JSON"), onchange)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onchange(onchange)
}

// Onchange fn
func (p *Props) Onchange(onchange func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('onchange', $1)", onchange)
	p.attrs["onchange"] = onchange
	return p
}

// Ondragstart fn
func Ondragstart(ondragstart func(e window.DragEvent)) *Props {
	macro.Rewrite("$1().Set('ondragstart', $2)", macro.Runtime("Map", "Set", "JSON"), ondragstart)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Ondragstart(ondragstart)
}

// Ondragstart fn
func (p *Props) Ondragstart(ondragstart func(e window.DragEvent)) *Props {
	macro.Rewrite("$_.Set('ondragstart', $1)", ondragstart)
	p.attrs["ondragstart"] = ondragstart
	return p
}

// Ondblclick fn
func Ondblclick(ondblclick func(e window.MouseEvent)) *Props {
	macro.Rewrite("$1().Set('ondblclick', $2)", macro.Runtime("Map", "Set", "JSON"), ondblclick)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Ondblclick(ondblclick)
}

// Ondblclick fn
func (p *Props) Ondblclick(ondblclick func(e window.MouseEvent)) *Props {
	macro.Rewrite("$_.Set('ondblclick', $1)", ondblclick)
	p.attrs["ondblclick"] = ondblclick
	return p
}

// Ondragover fn
func Ondragover(ondragover func(e window.DragEvent)) *Props {
	macro.Rewrite("$1().Set('ondragover', $2)", macro.Runtime("Map", "Set", "JSON"), ondragover)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Ondragover(ondragover)
}

// Ondragover fn
func (p *Props) Ondragover(ondragover func(e window.DragEvent)) *Props {
	macro.Rewrite("$_.Set('ondragover', $1)", ondragover)
	p.attrs["ondragover"] = ondragover
	return p
}

// Onkeydown fn
func Onkeydown(onkeydown func(e window.KeyboardEvent)) *Props {
	macro.Rewrite("$1().Set('onkeydown', $2)", macro.Runtime("Map", "Set", "JSON"), onkeydown)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onkeydown(onkeydown)
}

// Onkeydown fn
func (p *Props) Onkeydown(onkeydown func(e window.KeyboardEvent)) *Props {
	macro.Rewrite("$_.Set('onkeydown', $1)", onkeydown)
	p.attrs["onkeydown"] = onkeydown
	return p
}

// Onratechange fn
func Onratechange(onratechange func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('onratechange', $2)", macro.Runtime("Map", "Set", "JSON"), onratechange)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onratechange(onratechange)
}

// Onratechange fn
func (p *Props) Onratechange(onratechange func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('onratechange', $1)", onratechange)
	p.attrs["onratechange"] = onratechange
	return p
}

// Onscroll fn
func Onscroll(onscroll func(e window.UIEvent)) *Props {
	macro.Rewrite("$1().Set('onscroll', $2)", macro.Runtime("Map", "Set", "JSON"), onscroll)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onscroll(onscroll)
}

// Onscroll fn
func (p *Props) Onscroll(onscroll func(e window.UIEvent)) *Props {
	macro.Rewrite("$_.Set('onscroll', $1)", onscroll)
	p.attrs["onscroll"] = onscroll
	return p
}

// Onstalled fn
func Onstalled(onstalled func(e window.Event)) *Props {
	macro.Rewrite("$1().Set('onstalled', $2)", macro.Runtime("Map", "Set", "JSON"), onstalled)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Onstalled(onstalled)
}

// Onstalled fn
func (p *Props) Onstalled(onstalled func(e window.Event)) *Props {
	macro.Rewrite("$_.Set('onstalled', $1)", onstalled)
	p.attrs["onstalled"] = onstalled
	return p
}

// Attr fn
func Attr(key string, value interface{}) *Props {
	macro.Rewrite("$1().Set($2, $3)", macro.Runtime("Map", "Set", "JSON"), key, value)
	p := &Props{attrs: map[string]interface{}{}}
	return p.Attr(key, value)
}

// Attr fn
func (p *Props) Attr(key string, value interface{}) *Props {
	macro.Rewrite("$_.Set($1, $2)", key, value)
	p.attrs[key] = value
	return p
}