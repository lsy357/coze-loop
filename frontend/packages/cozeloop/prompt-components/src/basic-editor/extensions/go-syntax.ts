// Copyright (c) 2025 coze-dev Authors
// SPDX-License-Identifier: Apache-2.0
const tmLanguage = {
  foldingStartMarker: '\\{\\{\\s*(?:if|with|range)\\b',
  foldingStopMarker: '\\{\\{\\s*(?:else|end)\\b',
  name: 'go-template',
  patterns: [
    {
      begin: '\\{\\{',
      beginCaptures: {
        '0': {
          name: 'punctuation.section.embedded.begin.gotemplate',
        },
      },
      end: '\\}\\}',
      endCaptures: {
        '0': {
          name: 'punctuation.section.embedded.end.gotemplate',
        },
      },
      patterns: [
        {
          match: ':=',
          name: 'keyword.operator.initialize.gotemplate',
        },
        {
          match: '\\|',
          name: 'keyword.operator.pipe.gotemplate',
        },
        {
          match: '[.$][\\w]*',
          name: 'variable.other.gotemplate',
        },
        {
          match:
            '\\b(if|else|range|template|with|end|nil|with|define|block)\\b',
          name: 'keyword.control.gotemplate',
        },
        {
          match:
            '\\b(and|call|html|index|js|len|not|or|print|printf|println|urlquery|eq|ne|lt|le|gt|ge)\\b',
          name: 'support.function.builtin.gotemplate',
        },
        {
          begin: '/\\*',
          end: '\\*/',
          name: 'comment.block.gotemplate',
        },
        {
          begin: '"',
          beginCaptures: {
            '0': {
              name: 'punctuation.definition.string.begin.gotemplate',
            },
          },
          end: '"',
          endCaptures: {
            '0': {
              name: 'punctuation.definition.string.end.gotemplate',
            },
          },
          name: 'string.quoted.double.gotemplate',
          patterns: [
            {
              include: '#string_placeholder',
            },
            {
              include: '#string_escaped_char',
            },
          ],
        },
        {
          begin: '`',
          beginCaptures: {
            '0': {
              name: 'punctuation.definition.string.begin.gotemplate',
            },
          },
          end: '`',
          endCaptures: {
            '0': {
              name: 'punctuation.definition.string.end.gotemplate',
            },
          },
          name: 'string.quoted.raw.gotemplate',
          patterns: [
            {
              include: '#string_placeholder',
            },
          ],
        },
      ],
    },
  ],
  repository: {
    string_escaped_char: {
      patterns: [
        {
          match:
            '\\\\(\\\\|[abfnrtv\'"]|x[0-9a-fA-F]{2}|u[0-9a-fA-F]{4}|U[0-9a-fA-F]{8}|[0-7]{3})',
          name: 'constant.character.escape.gotemplate',
        },
        {
          match: '\\\\.',
          name: 'invalid.illegal.unknown-escape.gotemplate',
        },
      ],
    },
    string_placeholder: {
      patterns: [
        {
          match:
            "(?x)%\n                        (\\d+\\$)?                                    # field (argument #)\n                        [#0\\- +']*                                  # flags\n                        [,;:_]?                                     # separator character (AltiVec)\n                        ((-?\\d+)|\\*(-?\\d+\\$)?)?                     # minimum field width\n                        (\\.((-?\\d+)|\\*(-?\\d+\\$)?)?)?                # precision\n                        [diouxXDOUeEfFgGaAcCsSqpnvtTbyYhHmMzZ%]     # conversion type\n                    ",
          name: 'constant.other.placeholder.gotemplate',
        },
        {
          match: '%',
          name: 'invalid.illegal.placeholder.gotemplate',
        },
      ],
    },
  },
  scopeName: 'source.gotemplate',
  uuid: '43606de8-c638-11e2-b661-6711676f99ca',
};

export { tmLanguage };
