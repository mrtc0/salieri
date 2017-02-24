import React, { Component } from 'react';
import axios from 'axios';
import AceEditor from 'react-ace';
import { Header, Modal, Button, Menu, Segment, Grid, Icon, Message, Dropdown, Form, TextArea } from 'semantic-ui-react';

import 'brace/mode/c_cpp';
import 'brace/theme/monokai';
import 'brace/keybinding/emacs';
import 'brace/keybinding/vim';

// import { languageOptions } from '../common/language.js'
const languageOptions = [
    { key: 'clang', value: 'clang', text: 'clang' },
    { key: 'gcc', value: 'gcc', text: 'gcc' }
]
const KeyBindingOptions = [
    { key: 'vim', value: 'vim', text: 'vim' },
    { key: 'emacs', value: 'emacs', text: 'emacs' }
]

class App extends Component {
  constructor(props) {
    super(props);

    this.state= {
        language: "clang",
        result: "",
        stderr: "",
        stdin: "",
        code: "",
        activeItem: 'home',
        stdItem: 'stdout',
        keyBinding: '',
    };

    this.compile = this.compile.bind(this);
  }

  handleItemClick = (e, { name }) => this.setState({ activeItem: name })
  handleStdItemClick = (e, { name }) => this.setState({ stdItem: name })

  _onChange(newValue) {
    this.setState({code: newValue})
  }

  getStdin = (e, data) => this.setState({stdin: data.value})

  changeLanguage = (e, data) => this.setState({language: data.value})
  changeKeyBinding = (e, data) => this.setState({keyBinding: data.value})

  compile() {
    axios.post('http://localhost:8080/api/compiler/', {
        Code: this.state.code,
        Language: this.state.language,
        Stdin: this.state.stdin,
        Stdout: "",
        Stderr: ""
    }).then(response => {
        if (response.data.Stderr !== "") {
            this.setState({result: response.data.Stderr})
        } else {
            this.setState({result: response.data.Stdout})
        }
    })
    .catch(function (error) {
        console.log(error);
    })
  }

  componentDidMount() {
  }


  render() {
    const { activeItem, stdItem  } = this.state
    let stdfield;
    if (this.state.stdItem === "stdout") {
        stdfield = <Message><pre><code>{this.state.result}</code></pre></Message>
    } else {
        stdfield = <Form><TextArea placeholder='Stdin' onChange={this.getStdin} value={this.state.stdin}/></Form>
    }

    return (
      <div className="App">
          <Segment inverted>
              <Menu inverted pointing secondary>
                  <Menu.Item name='home' active={activeItem === 'home'} onClick={this.handleItemClick} />
                  <Menu.Menu position='right'>
                      <Dropdown item icon='wrench' simple>
                          <Dropdown.Menu>
                              <Modal trigger={<Dropdown.Item>Config</Dropdown.Item>}>
                                  <Modal.Header>Editor</Modal.Header>
                                  <Modal.Content>
                                      <Modal.Description>
                                          <Header>Keybinding</Header>
                                          <Dropdown placeholder='vim' search selection options={KeyBindingOptions} onChange={this.changeKeyBinding} />
                                      </Modal.Description>
                                  </Modal.Content>
                              </Modal>
                          </Dropdown.Menu>
                      </Dropdown>
                  </Menu.Menu>
              </Menu>
          </Segment>
          <Grid padded>
              <Grid.Row>
                  <Grid.Column width={8}>
                      <Dropdown placeholder='State' search selection options={languageOptions} onChange={this.changeLanguage} />
                      <Button primary onClick={this.compile}>
                          <Icon name="play" />
                          RUN
                      </Button>
                  </Grid.Column>
              </Grid.Row>
              <Grid.Row>
                  <Grid.Column width={8}>
                      <AceEditor
                          mode="c_cpp"
                          theme="monokai"
                          editorProps={{$blockScrolling: true}}
                          height="500"
                          width="800"
                          onChange={this._onChange.bind(this)}
                          value={this.state.code}
                          keyboardHandler={this.state.keyBinding}
                      />
                  </Grid.Column>
                  <Grid.Column width={8}>
                      <Menu tabular>
                          <Menu.Item name='stdout' active={stdItem === 'stdout'} onClick={this.handleStdItemClick} />
                          <Menu.Item name='stdin' active={stdItem === 'stdin'} onClick={this.handleStdItemClick} />
                      </Menu>
                      {stdfield}
                  </Grid.Column>
              </Grid.Row>
          </Grid>
      </div>
      );
      }
      }

      export default App;
