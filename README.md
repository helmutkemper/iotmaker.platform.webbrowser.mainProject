# iotmaker.platform.webbrowser.mainProject

todo: coisas demoradas/importantes devem ter um id por usuário para impedir o usupario de disparar o 
mesmo evento várias vezes seguidas. por exemplo, o usuário pode apertas o mesmo botão de gerar 
relatórios 500x seguidas e gerar 500 relatórios iguais.



Shape is the simplest display object you can add on stage. It is the most limited one: you can't add 
children to it (does not extend DisplayObjectContainer), does not have interactivity (does not extend 
InteractiveObject), does not have a timeline

Sprite extends DisplayObjectContainer and InteractiveObject, therefore it's interactive and you can 
add children to it. It's the most useful display class in my opinion, as long as you don't need a 
timeline.

MovieClip extends Sprite, so all of the above are true and you also get methods/properties associated 
with timeline control, but note that it's a dynamic class, so you can do some hacky thing on the fly, 
but you'll lose speed.
