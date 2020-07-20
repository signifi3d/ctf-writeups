
fd = open('blackmail_encrypted', 'rb')

search_str = "Isabelle"
alpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
remove_spice = lambda b: 0xff & ((b >> 1) | (b << 7))

line = fd.read()

for i in range(len(line)-8):
    poss_key = ""

    for j,c in enumerate(line[i:i+8]):
        for curr in alpha:
            if chr(remove_spice(c ^ ord(curr))) == search_str[j]:
                poss_key = poss_key + curr

    if len(poss_key) == 8:
        print(poss_key)
        
