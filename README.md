# criptodeslocamento
Criptografia usando deslocamentos


A função criptografar é responsável por aplicar a criptografia à mensagem original. Aqui está um passo a passo do processo:

A função recebe dois parâmetros: texto, que é a mensagem original que você deseja criptografar, e deslocamentos, que é uma sequência de dígitos que determina o deslocamento a ser aplicado a cada caractere do texto.

Ela começa por transformar a mensagem original em um slice de runes. Os runes são uma representação dos caracteres Unicode em Go, e essa conversão permite que cada caractere seja manipulado individualmente.

A função verifica se a sequência de deslocamentos não está vazia (ou seja, se possui comprimento maior que zero). Se estiver vazia, significa que não há deslocamento a ser aplicado, e a função retorna o texto original sem fazer nenhuma criptografia.

A seguir, a função itera pelos caracteres da mensagem original usando um loop for que percorre o slice de runes. Para cada caractere, ela determina o deslocamento a ser aplicado com base na sequência de deslocamentos. O deslocamento é calculado subtraindo o valor Unicode do caractere '0' (que é 48 em Unicode) do valor do dígito na sequência de deslocamentos. Isso converte o dígito de um caractere ASCII para um valor numérico.

O deslocamento calculado é então adicionado ao valor Unicode do caractere original, o que resulta em um novo caractere. Este novo caractere é armazenado de volta no slice de runes na mesma posição.

Enquanto a função itera pelos caracteres, ela também constrói a textoCriptografado concatenando cada caractere criptografado.

Após a iteração estar completa, a função utiliza a função base64.StdEncoding.EncodeToString([]byte(textoCriptografado)) para codificar a textoCriptografado em Base64. Isso é feito para que a mensagem criptografada possa ser representada em formato de texto seguro e legível, que pode ser transmitido ou armazenado facilmente.

Por fim, a função retorna a mensagem criptografada em Base64.

Resumindo, a criptografia funciona ao percorrer cada caractere da mensagem original, aplicar um deslocamento determinado pelos dígitos na sequência de deslocamentos, e em seguida, codificar a mensagem criptografada em Base64 para obter uma representação segura e legível da mensagem criptografada
