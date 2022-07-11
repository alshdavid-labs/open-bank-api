package pinpad

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/draw"
	imagePng "image/png"
)

var numbers = []*image.RGBA{
	b64ToRGBASure("iVBORw0KGgoAAAANSUhEUgAAALQAAABuCAYAAACOaDl7AAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsMAAA7DAcdvqGQAAAPhSURBVHhe7dyxShxRGIbh4xZ6Efa5kNyDkGpFkDRB06QNeAFJnyJdukDqpE0RLCxUSCXYBhSJhUK0cTLfOBNk8s/uOntWZj/eHx5017NbvRzOzu6amik20ur55mj3fDzaL39elQpgwK7qVnfVbp3x/VyM03r5x8PWA4DlULarhquYq525jvlya1TcbK8Udy9TUQADpkbVqpqtoz462Ulr6WxztNPETMhYNmq2iVotp/ocUtUePQAYOrVb79L72qFvdYPdGctK7dZB36Tql1K0EFgWTccEDQv9gt4L7gMGoF/QwEARNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQ9FN696IoPry6t/c8XoO5EPSiKd6fP4pw/lwXxcE34s6IoBflzbPukNujsL+8j58Hj0LQi6CYf53WtT5iiHpuBL0IOkb0HR1RoufETAg6NwXZNTqCfP1YFN8/F8Xvs/rO1uj+6HkxE4LO7fS4LrM17Z130rGEo0dvBJ2TLstFo105Wq+o9YKwPdrJo/WYiqBz0lEiGoUbrZeu83a0FlMRdE7REWLabvvpbb2wNbw47IWgc4qm67jR0O4dzbTHIUTQuXRd3Zhlp42GoHsh6Fy6gtYLxWj9Q9FRRVdLorWYiKBz0Y4aTbS2LbrUR9C9EHQuBD0IBJ0LQQ8CQedC0INA0LkQ9CAQdC7zBB19UEnvOkZrMRFB58J16EEg6Fz6Bt31gSa9JR6tx0QEnVM003ZafVQ0Gr5n2AtB5xS946f7orWN6HuHfMi/N4LOqevjo127re6PhheEvRF0Tl3nYe3S0Weio8t1Go4bvRF0bl2RKmqdl/UiUT+j44lGH/iPnhczIejcunbpWUZfx2J3ngtBL0LXlYtpw6W6uRH0ojwmav5zUjYEvUg6L3edqZvhf9tlRdBPQcHqOKE3WRqKfdK3wdELQcMKQcMKQcMKQcMKQcMKQcMKQcMKQcMKQcMKQcMKQcMKQcMKQcMKQcMKQcMKQcMKQcMKQcMKQcMKQcMKQcMKQcMKQcMKQcMKQcMKQcMKQcMKQcMKQcMKQcMKQcMKQcMKQcMKQcMKQcPKv6DPNke3+uUuWAQsA7VbB32Vzsejfd242V4JFwNDp3aroMuWdeTY1Y3LLXZpLB81q3aroMuW08lOWivLPmqiVu2EjaFTo2r1QczHajlpLsZpvYz6sP4DsFzKDVkNVzE3U2yk1fIPr8sXiQflouv/HgQMy3XVatms2r2vOKW/Wcl71ureRg0AAAAASUVORK5CYII="),
	b64ToRGBASure("iVBORw0KGgoAAAANSUhEUgAAALQAAABuCAYAAACOaDl7AAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsMAAA7DAcdvqGQAAAM2SURBVHhe7dyxattQFIfxGw3JQ2Tvg/QdAp1sAqFLsbt0LeQB2r1Dt26Fzu3ayUOGJGugD2ATmsGGOktu75Gl1jg3OHGkg/TnO/Ajlixp+rjIJnKoJx6F/dmwGM8GxST9nScR6LB51erY2q0yXs31IBymN883TgD6IbVrDZcxlytzFfPNcRGXJ3vx7nWIEegwa9RatWarqC+uRuEgTIfFqI6ZkNE31mwdtbUcqvuQsvbcCUDXWbvVKj2xFfrWNlid0VfWbhX0MpQvktyBQF/UHRM0JBA0pBA0pBA0pBA0pBA0pBA0pBA0pBA0pBA0pBA0pBA0pBA0pBA0pBA0pBA0pBA0pBA0pBA0pBA0pBA0pBA0pBA0pBA0pBA0pBA0pBA0pBA0pBA0pBA0pBA0pBA0pBC0pw+vYvz5NcZfl/+dvswfi50QtIdvH1fx5ubTm/w52AlBt8lC/j2tyn1gCLpRBN20dy9i/P55e8j1EHSjCLpJdo/8Z1GV+sgh6EYRdJPsA95Th6AbRdBNy334s3327UZuCLpRBN00C7Sesx+r25DN/etD0I0i6DZ8eX//+2WCdkHQXgjaBUF7IWgXBO2FoF0QtBeCdkHQXgjaBUF7IWgXBO2FoF0QtBeCdkHQXgjaBUF7IWgXBO2FoF0QtBeCdkHQXgjaBUF7IWgXBO2FoF0QdBvsn/ot1HX24GxubP/msfagbe662Iqg2/DUB2U3x37+IHddbEXQbXju2Kqduy62Iug2PHcIemcE3YbH/sjMQ2PPJOaui60IGlIIGlIIGlIIGvedZvb1BEFDCkFDCkFDCkFDCkFDCkFDCkFDCkFDCkFDCkFDCkFDCkFDCkFDCkFDCkFDCkFDCkFDCkFDCkFDCkFDCkFDCkFDCkFDCkFDCkFDCkFDCkFDCkFDCkFDCkFDCkFDCkFDyr+gp8Pi1l7cZQ4C+sDarYKeh9mgmNjG8mQvezDQddZuGXRq2W45xrZxc8wqjf6xZq3dMujUcrgahYNU9kUdtdVO2Og6a9RaXYv50loONteDcJiiPq/eAPolLcjWcBlzPfEo7Kc33qYPiWfpoMW9k4BuWZStpmat3VXFIfwFjkXTNh4DxL4AAAAASUVORK5CYII="),
	b64ToRGBASure("iVBORw0KGgoAAAANSUhEUgAAALQAAABuCAYAAACOaDl7AAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsMAAA7DAcdvqGQAAAOuSURBVHhe7dsxS9xgHMfxxxv0Rbj3hfQ9CJ1OBOlStEvXgi+g3Tt061bo3K4dioODCp0E14IidVCoLj7NL5eUIz6hiUkudz++f/hQr5fL9PXhSS6GcuJWWL/cnuxfTieH2b83mQgssZui1X21W2Q8m6tp2MzePK58ACOJByH5/6iRtauG85jzlbmI+XpnEu921+LDyxAjsMTUqFpVs0XUJ2d7YSNcbE/2ypgJGatGzZZRq+VQ7EPy2lMfAJad2i1W6UOt0Pd6weqMVaV2i6DvQv5DJnUgsCrKjgkaFggaVggaVggaVggaVggaVggaVggaVggaVggaVggaVggaVggaVggaVggaVggaVggaVggaVggaVggaVggaVggaVggaVggaVggaVggaVggaVggaVggaVggaVggaVgh6Ud69iPHDqxn9nDoGnRH0UN48i/HL+xjPT2Pt/PwxCzz1eTwJQQ/h68cY/9wW1TaYo2+zX4DUudAKQfdJUf46LyptOVqtU+dEKwTdp4PnRZ1PHG1RUudFYwTdN6201dE+WtsQSb1fzu+L9DnRGEH3Tat0uX9WwHpdPUZ3Oer22NwB6YSgh6A7F6mQ52l7kRr9EqSORyMEPRZdQKaGoDsh6DGlhqA7IeixsEIPgqDH8ultUXBluCjshKDHom8Hq6M7H6lj0RhBj6HuCxi2G50R9BhSDyxpdeZ5js4IetG0CqeG1bkXBL1Idd8Q6oGm1PFojaAXRdsJPatRHQXOnY3eEPSi1D3ozxN2vSLoRfj+uai3Mrp1lzoeT0bQQ6t7CEkrdup4dELQQ9LeODW6COQW3SAIeijzz0XPD/ebB0XQQ1Cwqb8t5I7G4Ah6CHV/ZkXMgyPovtV9E6jVWReC/6M7IqnzohGC7lPdRWCbUdSpc6MRgu5T3S26NkPQnRB0n+q2G22GoDsh6D6xQo+OoGGFoGGFoGGFoGGFoGGFoGGFoGGFoGGFoGGFoGGFoGGFoGGFoGGFoGGFoGGFoGGFoGGFoGGFoGGFoGGFoGGFoGGFoGGFoGGFoGGFoGGFoGGFoGGFoGGFoGGFoGGFoGGFoGGFoGGFoGGFoGHlX9AX25N7/fCQOAhYBWq3CPomXE4nh3pxt7uWPBhYdmo3DzprWVuOfb243mGVxupRs2o3DzprOZzthY2s7JMyatVO2Fh2alStzsV8qpaD5moaNrOoj4s3gNWSLchqOI+5nLgV1rM3XmcXiUfZQbePPgQsl9u81axZtTurOIS/N+H6tIqgtZMAAAAASUVORK5CYII="),
	b64ToRGBASure("iVBORw0KGgoAAAANSUhEUgAAALQAAABuCAYAAACOaDl7AAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsMAAA7DAcdvqGQAAAPWSURBVHhe7dyxShxBAMbx8Qp9CPs8SN5BSHUihDRB06QN+ABJnyJdukDqpE0RLCxUSCXYBhSJhUK0cbLfupscl1kd3V137+M/8EPPm93qzzC3u2eoR1wLyyfrk62T6WSn+HleiMCInVetbqndKuObcToNq8Wbe3MHAIuhaFcNlzGXK3MV89nGJF4+X4rXL0KMwIipUbWqZquo9w83w0o4Xp9s1jETMhaNmq2jVsuh2oeUtacOAMZO7Var9I5W6Cu9YHXGolK7VdCXofylkJoILIq6Y4KGBYKGFYKGFYKGFYKGFYKGFYKGFYKGFYKGFYKGFYKGFYKGFYKGFYKGFYKGFYKGFYKGFYKGFYKGFYKGFYKGFYKGFYKGFYKGFYKGFYKGFYKGFYKGFYKGFYKGFYKGFYJ+DK+fxPj+5T/bT9Pz0BpB9+Xtsxi/fYrx13FMjt8XMe5+Je6OEXTXFOjRQVVt5vj8Ln0u3BtBd+3Lh6rSew4dlzof7oWgu6YV+qFD25TUOZGNoPugvXE9fny/2VLUHwi1r9b+OTU0N3U+ZCPoPuiqhrYQ+pl6XytxU9Sp+chG0ENp2mtrFU/NRxaCHorCTQ2CboWgh0LQvSDooejD4fzQvjo1F9kIeghNHwoVeWo+shH0Y9J2oumy3c+j5qsiyEbQffr4pqr1jqGYeaajEwTdp5zb4LoJkzoWD0LQfcp9rkNRs93oBEH3Sbe8c4f21TzL0RpB92n+wX6t2FqNm257ay+dOg+yEfQQFLriTQ2ejW6FoIeiqFMrtb4ckJqPLAQ9pNnHTGdHai6yEPSQmq6CpOYiC0EPKfU8h0ZqLrIQ9FC0h059I5wrHa0QdNe0L5bbbmUr5qb9s/6eOgZZCLpLutY8O7Taap88fy266X91aGhO6tzIQtBd0pdc2wwu2bVG0F1qM3RNWluR1HmRjaC7pO1E023t2waPj3aGoLumVfaufXI9FDK3ujtF0H3Sqlt/EJylv7Ei94KgYYWgYYWgYYWgYYWgx2A78Tc8CEHDCkHDCkHDCkHDCkHDCkHDCkHDCkHDCkHDCkHDCkHDCkHDCkHDCkHDCkHDCkHDCkHDCkHDCkHDCkHDCkHDCkHDCkHDCkHDCkHDCkHDCkHDCkHDCkHDCkHDyt+gj9cnV/rlOjEJWARqtwr6PJxMJzt6cfl8KTkZGDu1WwZdtKwtx5ZenG2wSmPxqFm1WwZdtBwON8NKUfZ+HbVqJ2yMnRpVqzMxH6jloHE6DatF1HvVG8BiKRZkNVzGXI+4FpaLN14VHxJ3i0kX/x0EjMtF2WrRrNq9qTiEP7EBBoSlyYaUAAAAAElFTkSuQmCC"),
	b64ToRGBASure("iVBORw0KGgoAAAANSUhEUgAAALQAAABuCAYAAACOaDl7AAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsMAAA7DAcdvqGQAAAN6SURBVHhe7dsxTxRBGIfx4Qr4EPR+EL8DidUREmJjOBtbEz6A9hZ2dibW2loYCgogsSKhNYEQKSARGtZ5l1lzLC8Gb2bZ3b/Pm/wix83F5nGys7eGZqq1sHyyPpmdTCc78c/zqAIG7Dy1OrN2U8Y3czoNq/HNvdYHgHGI7VrDdcz1zpxiPtuYVJebS9X181BVwIBZo9aqNZui3j/cCivheH2y1cRMyBgba7aJ2loO6Tqkrt37ADB01m7apXdsh76yF+zOGCtrNwV9GeofIm8hMBZNxwQNCQQNKQQNKQQNKQQNKQQNKQQNKQQNKQQNKQQNKQQNKQQNKQQNKQQNKQQNKQQNKQQNKQQNKQQNKQQNKQQNKQQNKQQNKQQNKQQNKQQNKQQNKQQNKQQNKQQNKQQNKQQNKQT9WHa/VNXRwW3vXvhrsTCCfgwfXlfubD/112NhBP0Yvn9LBc+N7dDeWmQh6K7ZLuzNp7f+emQh6K59fp8KnptfF/5aZCPorlm87bHIvbXIRtBdsssKbzgMdoagu2QHv/bYAdFbiyIIuiv3HQa599wpgu7K14+p4Ln5eeyvRTEE3RUOg70g6C7cdxh89cRfj2IIugveYdCe5ZhfY1+H247dsH8E3P3IRtClvXmWCm6N/X5+nTdckmQj6NJsJ27Pj6O767wh6GwEXZJdI3uHQe+5DW8IOhtBl+QdBi1w7zDoTU7Q287v/kMEXZJdWrTH7kd7a71hh85G0KXcdxhcdHgibyEEXYrtrqXH+3vwVwRdCkEPAkGXQtCDQNCl2Dd/zf/mfghv7OGl5v32N4t4EILuizfc5chG0H3xhqCzEXRfvCHobATdF28IOhtB98Ubgs5G0H2xOxrt8R5iwj8haEghaEghaEghaEghaEghaEghaEghaEghaEghaEghaEghaEghaEghaEghaEghaEghaEghaEghaEghaEghaEghaEghaEghaEghaEghaEghaEghaEghaEghaEghaEghaEghaEghaEghaEghaEj5E/Tx+uTKfrh2FgFjYO2moM/DyXSyYy8uN5fcxcDQWbt10LFlu+SY2YuzDXZpjI81a+3WQceWw+FWWIll7zdRW+2EjaGzRq3VuZgPrOVgczoNqzHqvfQGMC5xQ7aG65ibqdbCcnzjZTwk7sZFF3c+BAzLRd1qbNbavak4hN/IbvBEIFHtJwAAAABJRU5ErkJggg=="),
	b64ToRGBASure("iVBORw0KGgoAAAANSUhEUgAAALQAAABuCAYAAACOaDl7AAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsMAAA7DAcdvqGQAAAOESURBVHhe7dqxTttQGIbhQwa4CPZeSO8BqVMQEupSQZeulbiAdu/QrVulzu3KlIEBWJFYK4FQGUAqLLjnc+wqoifgOD51/On9pUeNie3p7ZFjO9RTbIX1y+3R/uV4NIn/3kQFsMJuqlb31W6V8XSuxmEzfnn86ABgGGK7ariMuVyZq5ivd0bF3e5a8fA6FAWwwtSoWlWzVdQnZ3thI1xsj/bqmAkZQ6Nm66jVcqiuQ8raUwcAq07tVqv0RCv0vTZYnTFUarcK+i6UH6LUjsBQ1B0TNCwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNPI5SPwtM4KGFYKGFYKGFYKGFYKGFYKGFYKGFYKGFYLO4fvnojg/befdi/Q50QhB56Aw286nN+lzohGCzoGge0PQORB0bwg6h2WCPniZPicaIegcUkGz8v4XBJ0DQfeGoHMg6N4QdA4E3RuCzoGge0PQOaSCPvoxfYJYI/AsCDqHVNDzRqFzq64zBJ3DIkFrft8WxbeP6XNhIQSdw6JB1/Plffp8aIygc5gN+uf5dFv0+anRSs3bdksh6Bx0TTzvuljB6keh4k2Nvksdh0YIui8fXlUFP5pfF+n90QhB90l3OFLDZUdrBN0n3dlIDfeoWyPoPinc1BB0awTdJ4LuHEH3SXc0UsM1dGsE3RdFqzsaj4e7HEsh6K4dfp0+RHnq/QzFPO8Oh45PHYNGCLpLinh29GSwfrOupu3UylwPLyothaC7NG/VbTo8JVwaQXfpqZX3udF/htQ5sRCC7pIuKdpEzXVzZwg6Bz0BfO7NOo1WZb3TkToHWiHonPQDr/4hOEt/S+2PpRE0rBA0rBA0rBA0rBA0rBA0rBA0rBA0rBA0rBA0rBA0rBA0rBA0rBA0rBA0rBA0rBA0rBA0rBA0rBA0rBA0rBA0rBA0rBA0rBA0rBA0rBA0rBA0rBA0rBA0rBA0rBA0rBA0rBA0rBA0rPwN+mJ7dK8PD4mdgCFQu1XQN+FyPJpo4253LbkzsOrUbhl0bFmXHPvauN5hlcbwqFm1WwYdWw5ne2Ejln1SR63aCRurTo2q1ZmYT9Vy0FyNw2aM+rj6AhiWuCCr4TLmeoqtsB6/eBt/JB7FnW7/OQhYLbdlq7FZtTutOIQ/C+oaohBerC0AAAAASUVORK5CYII="),
	b64ToRGBASure("iVBORw0KGgoAAAANSUhEUgAAALQAAABuCAYAAACOaDl7AAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsMAAA7DAcdvqGQAAAPzSURBVHhe7duxThRdHIbxwxZwEfReiPdAYrWEhNgYsLE14QK0t7CzM7HW1sJQUACJFQmtCYRIAYnQMM67zOhm/c+yO3NGd1+fk/zyLbtnt3q+kzNnxlSPYiOtnm8Ods+Hg/3yv1elAlhgV1Wru2q3yvh+XAzTevnh4cQXgOVQtquGRzGPVuYq5sutQXGzvVLcPU1FASwwNapW1WwV9dHJTlpLZ5uDnTpmQsayUbN11Go5VfuQUe3RF4BFp3arVXpfK/St/mB1xrJSu1XQN2n0ohRNBJZF3TFBwwJBwwpBwwpBwwpBwwpBwwpBwwpBwwpBwwpBwwpBwwpBwwpBwwpBwwpBwwpBwwpBwwpBwwpBwwpBwwpBwwpBwwpBwwpBwwpBwwpBwwpBwwpBwwpBwwpBwwpBwwpB/wtvnv324lE8B60Q9N/y4XVRfDstwvH9rCg+viXuDAi6b6+eNIc8OX5c36/a0e9gJgTdJ8WsSOcZWqmj38JMCLovbWLWfLYdnRB0X5q2GXpfq3B9UajX9VxW584Iug+6AIzGwad4vmhFj97HXAi6Dzq1mBxahaO5yIqgc3v3sip4Yuj9aD6yIujctK2YHLrYi+YiO4LOLTrZmLZ3RlYEnZMu7KIxfnpRn2zUuJGSFUHnNG3/rHibzqX1vj7nDLozgs5JUUYjOvWIhk5CiLoTgs7p8/uqzA6D471OCDqn0+Oqyo5DK330+3gQQefUFHT9eOj43UC9/vqlmjAxNH/8dzEzgs4pCvqhLUTT/wTcCm+FoHOK4tR70dxa08mIngeJ5mMqgs6pTdASDfbRrRB0TtEpxyz74WgQdCsEnVPTOXQ0d1w0CLoVgs5Jt7GjMe1Ju6bvcEu8FYLOSXf5ojHt4aTo6TwN7hi2QtC5zXMM1/QwE0/ntUbQuTX98ys9gFQfxWn11eumh5XYbrRG0H1oWqVnGTopiX4TMyHoPmgr0bT6Ths8mNQZQfdl3qi1b+ZCsDOC7pMCVajTwtaqzJ45G4L+WxStbpbUdFG49ziei9YIGlYIGlYIGlYIGlYIGlYIGlYIGlYIGlYIGlYIGlYIGlYIGlYIGlYIGlYIGlYIGlYIGlYIGlYIGlYIGlYIGlYIGlYIGlYIGlYIGlYIGlYIGlYIGlYIGlb8g94L3oMt/6DxXyFoWCFoWCFoWCFoWPkV9Nnm4FYv7oJJwDJQu1XQV+l8ONjXHzfbK+FkYNGp3VHQZcvacuzqj8stVmksHzWrdkdBly2nk520VpZ9VEet2gkbi06NqtWxmI/VctK4GKb1MurD6gNguZQLshoexVyPYiOtlh88Ly8SD8pJ1398CVgs16NWy2bV7n3FKf0EZsBlsnrV+1AAAAAASUVORK5CYII="),
	b64ToRGBASure("iVBORw0KGgoAAAANSUhEUgAAALQAAABuCAYAAACOaDl7AAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsMAAA7DAcdvqGQAAANzSURBVHhe7dsxS1tRGIfxYwb9EO79IP0OQqcEQbqUpEvXgh+g3Tt061bo3K4dioODCp0E10JE6qDQuHh73utJCekbNeace3P/PC/80JgTXR4ON7nHMJ1qJ2yeD3qj837vIH69iipgjV2lVkfWbsr4bi76YTs+eTT3AqAbYrvWcB1zvTOnmC93e9Vkb6O6fRmqClhj1qi1as2mqI9Ph2ErjAe94TRmQkbXWLPTqK3lkK5D6tq9FwDrztpNu/SB7dA39oDdGV1l7aagJ6H+JvIWAl0x7ZigIYGgIYWgIYWgIYWgIYWgIYWgIYWgIYWgIYWgIYWgIYWgIYWgIYWgIYWgIYWgIYWgIYWgIYWgIYWgIYWgIYWgIYWgIYWgIYWgIYWgIYWgIYWgIYWgIYWgIYWgIYWgIYWgc9p/XlVnJ6v7+aOq3jzz/wbuRdA5fXhVZRv7Xd7fwL0IOieCbh1B55QzaC45noSgc3r3ItW44vw6838/HkTQbfs9ThXPzJf3/lo8iKDb9OltKnhm/lxzubECgm6TfUQ3P4ff/LV4FIJuy6Lrbfu5tx6PQtBtsZ14fmzH9tbi0Qi6DXaN7A1vBldG0G34+jEVPDP2ZtBbi6UQdBss3vmxyL21WApBN80uK7yxg03eeiyFoJvm3Uix03XeWiyNoJu06KwHB5GyIegmeTdSbMf21uJJCLopdo3sDW8GsyLopng3Uji3kR1BN8Gi9T6q49xGdgTdBO9Gig3nNrIj6CZ4H9VxiL8Igi5t0Y0Uzm0UQdCl2U48P5zbKIagS1p0I+X7Z389VkbQJdktbW84t1EMQZey6EYK5zaKIuhS7LLCG/vHWG89siDoUryP6vgXq+IIGlIIGlIIGlIIGlIIGlIIGlIIGlIIGlIIGlIIGlIIGlIIGlIIGlIIGlIIGlIIGlIIGlIIGlIIGlIIGlIIGlIIGlIIGs3Yd35WAEFDCkFDCkFDCkFDCkFDCkFDCkFDCkFDCkFDCkFDCkFDCkFDCkFDyr+gx4PejX1z6ywCusDaTUFfhfN+78AeTPY23MXAurN266Bjy3bJMbIHl7vs0ugea9barYOOLYfTYdiKZR9Po7baCRvrzhq1VmdiPrGWg81FP2zHqI/SE0C3xA3ZGq5jnk61EzbjE6/jm8TDuOj6vxcB6+W6bjU2a+3eVRzCX0l9toRFpJLFAAAAAElFTkSuQmCC"),
	b64ToRGBASure("iVBORw0KGgoAAAANSUhEUgAAALQAAABuCAYAAACOaDl7AAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsMAAA7DAcdvqGQAAAQNSURBVHhe7duxShxRGIbhcQu9CPtcSO5BSLUiSJqgadIGvICkT5EuXSB10qYIFhYqpBJsA4rEQiHaOJlvnUmW9T/u7OwZ2f14f3iIurNTvRzOntkUzZQbxer55mD3fDjYr/69qpTAAruqW91Vu3XG93MxLNarFw8n3gAsh6pdNTyKebQy1zFfbg3Km+2V8u5lUZbAAlOjalXN1lEfnewUa8XZ5mCniZmQsWzUbBO1Wi7qfcio9ugNwKJTu/Uqva8V+la/sDpjWandOuibYvRDJboQWBZNxwQNCwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwQNKwT9lPael+WHV/9F12AuBN03RXzwrSx/n5Xh/Doty68fy/LNs/j9mAlB90mhth0F/+5FfB+0RtB90ao86/y5Juo5EXQfPr2tC+0w2oJE90QrBN2H1H5ZW5DmGkWvFTkavTZ+P7RG0LlpyxDNeMzTrv3++eG1aIWgc0t9EEydYmiLMTmnx/G1mIqgc0sFHV0rindyCLozgs4tFXRqXxzto9lydEbQuaVOOLS1mNx2pOLnKWJnBJ2bngymRqcfTaz6QBitzj9/PLwnWiPoPkx7qKI9chRztIpjJgTdB63SqTPm1GhlJua5EXRftLVoGzUrczYE3SedVrQdxc+HwbkRdB+02kbny23my/v4nmiFoHNTzNHTP/1NR3ptQucbd50RdG5RsJN7ZG0tHgtbr43fE60RdE7aLkSjU4/o+tSDFU3qPXgUQeeko7fJmbbapqJmL90JQecUHdMp2OjahrYi0Ux7H0IEnVM0bcKMhqA7Ieicopm25dAHxGgIuhOCzik6rtOk9sOpIz4NR3edEHROWlVTo5VaYWtFFj1FTD0a17fyovtjKoLOSSvurF9Kiob/JNsZQeeW+oJ/29FXT6P7ohWC7oOi7rJS80FwbgTdF20/FOi0sPW6VmWeDGZB0E9BJxb6QKjAG/qdk4zsCBpWCBpWCBpWCBpWCBpWCBpWCBpWCBpWCBpWCBpWCBpWCBpWCBpWCBpWCBqLbS/42yMIGlYIGlYIGlYIGlYIGlYIGlYIGlYIGlYIGlYIGlYIGlYIGlYIGlYIGlYIGlYIGlYIGlYIGlYIGlYIGlb+BX22ObjVD3fBRcAyULt10FfF+XCwr19utlfCi4FFp3ZHQVcta8uxq18ut1ilsXzUrNodBV21XJzsFGtV2UdN1KqdsLHo1KhaHYv5WC0XmothsV5FfVi/ACyXakFWw6OYmyk3itXqhdfVh8SD6qLrB28CFsv1qNWqWbV7X3FR/AXS3qnLffQxBAAAAABJRU5ErkJggg=="),
	b64ToRGBASure("iVBORw0KGgoAAAANSUhEUgAAALQAAABuCAYAAACOaDl7AAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsMAAA7DAcdvqGQAAAP6SURBVHhe7dy9ThRRHIbxwxZwEfReiPdAYrWEhNgYsLE14QK0t7CzM7HW1sJQUACJFQmtCYRIAYnQMM67zChZ/7OzzJ5xd988J/kFlj2z1ePJmY811aPYSKvnm4Pd8+Fgv/x5VSqABXZVtbqrdquM78fFMK2Xbx6OHQAsh7JdNTyKebQyVzFfbg2Km+2V4u55KgpggalRtapmq6iPTnbSWjrbHOzUMRMylo2araNWy6nah4xqjw4AFp3arVbpfa3Qt3rB6oxlpXaroG/S6JdSNBFYFnXHBA0LBA0rBA0rBA0rBA0rBA0rBA0rBA0rBA0rBA0rBA0rBA0rBA0rBA0rBA0rBA0rBA0rBA0rBA0rBA0rBA0rBA0rBA0rBA0rBN1kL/gbFh5BwwpBwwpBwwpBwwpBwwpBwwpB/y+vnhTFuxd/6XU0DzMh6D4p2s/vi+LnWREO/V3vE3c2BN2XD6+L4td1VW7L0LxPb+PPwaMQdB8UZ5dB1DMj6NzePKvq7Dh0fPS5mApB53Z6XJU5Nn6c3q/AOiHUT72Oho6PPhdTIeicFGs0vn+L5zfFzyrdGUHn9PVjVeTYaLqKsfe0mjA2dOUjmo9WBJ1TtI1oWp1r0TFsOzoj6Jyi0bbaNq3q0Vy0IuicotEWtN6PRjQXrQg6p2gcfInn1ppOJPX3aD4mIuicov2w7gJOurVN0FkRdE5N+2Gt0lHUusqhk8ZoEHQnBJ1T02U4Da3UClt7ZoXfdGOlHgTdCUHn1rRKP3YQdCcE3QetxLMO7hZ2QtB90dai7fFRhd/0ZF70mWhF0H3SiaCCVbi6+yc6CVTs2m9rTnQdWv8Qxj8LUyHoeYv23Nz67oyg5y164k6RR3PRiqDnSVuSaOjrW9F8tCLoeYpOCNk/z4Sg5yn6NrhOIKO5mApBz0vTtWpuqMyEoPugKDUUrfbDD5/j0OvoRFCj7csAaEXQfWj6j2UmDe2d62vT6Iygc9MK3GVwZSMLgs7tsUFrZWbfnA1B90FbB+2fJz3Lofc0Z9LD/3g0gu6bVl89r/EQK3JvCBpWCBpWCBpWCBpWCBpWCBpWCBpWCBpWCBpWCBpWCBpWCBpWCBpWCBpWCBpWCBpWCBpWCBpWCBpWCBpWCBpWCBpWCBpWCBpWCBpWCBpWCBpWCBpWCBpWCBpWCBpWCBpWCBpWCBpWCBpW/gR9tjm41S93wSRgGajdKuirdD4c7OvFzfZKOBlYdGp3FHTZsrYcu3pxucUqjeWjZtXuKOiy5XSyk9bKso/qqFU7YWPRqVG1+iDmY7WcNC6Gab2M+rB6A1gu5YKshkcx16PYSKvlGy/Lk8SDctL1PwcBi+V61GrZrNq9rzil38xKYHZ08qFRAAAAAElFTkSuQmCC"),
}

func b64ToRGBASure(s string) *image.RGBA {
	v, _ := b64ToImage(s)
	return imageToRGBA(v)
}

func b64ToImage(s string) (image.Image, error) {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	img, err := imagePng.Decode(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	return img, nil
}

func imageToRGBA(src image.Image) *image.RGBA {
	if dst, ok := src.(*image.RGBA); ok {
		return dst
	}
	b := src.Bounds()
	dst := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(dst, dst.Bounds(), src, b.Min, draw.Src)
	return dst
}