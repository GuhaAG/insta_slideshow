import React from 'react'

class Demo extends React.Component {
    constructor() {
        super();
        this.state = { images: [] };
    }

    componentDidMount() {
        fetch('api/v1/images').then((response) => response.json()).then((js) => this.setState({ images: js }));
    }

    render() {
        let images = this.state.images;
        return (
            <ul>
                {images.map(img =>
                    <li>{img}</li>
                )}
            </ul>
        );
    }

}

export default Demo;